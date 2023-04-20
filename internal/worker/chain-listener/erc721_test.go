package chainListener

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/phhphc/nft-marketplace-back-end/internal/contracts"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
	"github.com/stretchr/testify/suite"
)

type NftTestSuite struct {
	suite.Suite
	auth    *bind.TransactOpts
	address common.Address
	gAlloc  core.GenesisAlloc
	sim     *backends.SimulatedBackend
	gasLim  uint64
	erc721  *contracts.ERC721
	lg      log.Logger
}

func TestRunNftTestSuite(t *testing.T) {
	suite.Run(t, new(NftTestSuite))
}

// SetupTest is called before each test
// Step 1: Set up the private key, transaction, and simulated backend
func (s *NftTestSuite) SetupTest() {
	fmt.Println("Start to setup...")
	s.lg = *log.GetLogger()
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("fail in generate private key")
	}

	s.auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("fail in create transaction")
	}

	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10)

	s.address = s.auth.From
	s.gAlloc = map[common.Address]core.GenesisAccount{
		s.address: {Balance: balance},
	}

	s.gasLim = uint64(4712388)
	s.sim = backends.NewSimulatedBackend(s.gAlloc, s.gasLim)

	// Load the contract ABI and bytecode
	abi, err := abi.JSON(strings.NewReader(contracts.ERC721MetaData.ABI))
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("fail in load erc721 abi")
	}
	bytecode, err := hexutil.Decode(contracts.ERC721MetaData.Bin)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("fail in load erc721 bytecode")
	}

	address, tx, _, err := bind.DeployContract(s.auth, abi, bytecode, s.sim)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("fail in deploy contract")
	}

	s.erc721, err = contracts.NewERC721(address, s.sim)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("fail in create new contract instance")
	}

	s.Nil(err)
	s.sim.Commit()

	_ = tx

}

func (s *NftTestSuite) SetupAccounts() {

}

func (s *NftTestSuite) TestMint() {
	transactor := s.erc721.ERC721Transactor
	_ = transactor
}
