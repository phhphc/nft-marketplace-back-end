package services

import (
	"context"
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang-jwt/jwt"
	"github.com/phhphc/nft-marketplace-back-end/configs"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
	"github.com/spruceid/siwe-go"
	"math/big"
	"regexp"
	"strconv"
	"time"
)

type AuthenticationService interface {
	GetUserNonce(ctx context.Context, address string) (string, error)
	Login(ctx context.Context, address string, messageStr string, sigHex string) (string, error)
	GenerateJWTToken(user *entities.User) (string, error)
	ValidateJWTToken(tokenString string) (jwt.MapClaims, error)
}

const (
	ADMIN_ROLE_ID     = 1
	MODERATOR_ROLE_ID = 2
	USER_ROLE_ID      = 3
)

func (s *Services) isValidAddress(address string) bool {
	if len(address) != 42 {
		return false
	}

	match, _ := regexp.MatchString("^0x[0-9a-fA-F]{40}$", address)
	return match
}

func (s *Services) GetUserNonce(ctx context.Context, address string) (string, error) {
	// Check if the user is in the database
	etherAddress := common.HexToAddress(address)
	res, err := s.repo.GetUserByAddress(ctx, etherAddress.Hex())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) && s.isValidAddress(etherAddress.Hex()) {
			nonce := s.generateNonce()
			user, err := s.repo.InsertUser(ctx, postgresql.InsertUserParams{
				Nonce:         nonce,
				PublicAddress: etherAddress.Hex(),
			})
			if err != nil {
				return "", err
			}

			_, err = s.repo.InsertUserRole(ctx, postgresql.InsertUserRoleParams{
				RoleID:  USER_ROLE_ID,
				Address: etherAddress.Hex(),
			})

			if err != nil {
				return "", err
			}
			return user.Nonce, nil
		}
		return "", err
	}
	return res.Nonce, nil
}

func (s *Services) verifySignature(from string, sigHex string, message []byte) (bool, error) {
	sig := hexutil.MustDecode(sigHex)

	msg := accounts.TextHash(message)
	sig[crypto.RecoveryIDOffset] -= 27

	pubKey, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false, err
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return from == recoveredAddr.Hex(), nil
}

func (s *Services) updateUserNonce(ctx context.Context, address string, nonce string) (*entities.User, error) {
	// Check if the user is in the database
	arg := postgresql.UpdateNonceParams{
		Nonce:         nonce,
		PublicAddress: address,
	}
	res, err := s.repo.UpdateNonce(ctx, arg)
	if err != nil {
		return nil, err
	}
	user := entities.User{
		Address: res.PublicAddress,
		Nonce:   res.Nonce,
	}
	return &user, nil
}

func (s *Services) generateNonce() string {
	nonceBytes := make([]byte, 32)

	_, err := rand.Read(nonceBytes)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	nonce := new(big.Int).SetBytes(nonceBytes).String()
	fmt.Printf("nonce: %v", nonce)
	return nonce
}

func (s *Services) Login(ctx context.Context, address string, messageStr string, sigHex string) (string, error) {
	// Check if the user is in the database
	etherAddress := common.HexToAddress(address)

	res, err := s.repo.GetUserByAddress(ctx, etherAddress.Hex())
	if err != nil {
		return "", err
	}

	fmt.Printf("Address %s\n", etherAddress.Hex())

	var message *siwe.Message
	message, err = siwe.ParseMessage(messageStr)
	if err != nil {
		return "", err
	}

	// Check is the nonce in the message
	if message.GetNonce() != res.Nonce {
		return "", fmt.Errorf("invalid nonce")
	}

	// Verify the signature
	isValid, err := s.verifySignature(etherAddress.Hex(), sigHex, []byte(messageStr))
	if err != nil {
		return "", err
	}
	if !isValid {
		return "", fmt.Errorf("invalid signature")
	}

	// Update the nonce
	nonce := s.generateNonce()
	user, err := s.updateUserNonce(ctx, etherAddress.Hex(), nonce)
	if err != nil {
		return "", err
	}

	// Get the JWT token
	token, err := s.GenerateJWTToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Services) GenerateJWTToken(user *entities.User) (string, error) {
	cfg, err := configs.LoadConfig()
	if err != nil {
		return "", err
	}
	fmt.Printf("secret: %s\n", cfg.JwtSecret)
	secret := []byte(cfg.JwtSecret)
	jwtExpired, err := strconv.Atoi(cfg.JwtExpired)
	if err != nil {
		return "", err
	}
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["address"] = user.Address
	claims["nonce"] = user.Nonce // current nonce for protect replay attack
	claims["exp"] = time.Now().Add(time.Duration(jwtExpired) * time.Second).Unix()
	claims["roles"] = user.Roles

	tokenString, err := token.SignedString(secret)
	fmt.Printf("token: %s\n", tokenString)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *Services) ValidateJWTToken(tokenString string) (jwt.MapClaims, error) {
	cfg, err := configs.LoadConfig()
	if err != nil {
		return nil, err
	}
	secret := []byte(cfg.JwtSecret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
