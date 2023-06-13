package chainListener

import (
	"context"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/contracts"
	"github.com/phhphc/nft-marketplace-back-end/internal/services"
	"github.com/phhphc/nft-marketplace-back-end/pkg/clients"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
)

type ChainListener interface {
	Run(ctx context.Context) error
}

type worker struct {
	ethClient *clients.EthClient
	lg        *log.Logger
	mkpAddr   common.Address

	erc721Abi      abi.ABI
	erc721Contract *contracts.ERC721Filterer
	mkpAbi         abi.ABI
	mkpContract    *contracts.MarketplaceFilterer

	Service services.Servicer
}

func NewChainListener(service services.Servicer, ethClient *clients.EthClient, mkpContractAddr string) (ChainListener, error) {
	mkpAddr := common.HexToAddress(mkpContractAddr)
	mkpAbi, err := abi.JSON(strings.NewReader(contracts.MarketplaceMetaData.ABI))
	if err != nil {
		return nil, err
	}
	mkpContract, err := contracts.NewMarketplaceFilterer(mkpAddr, nil)
	if err != nil {
		return nil, err
	}

	erc721Abi, err := abi.JSON(strings.NewReader(contracts.ERC721MetaData.ABI))
	if err != nil {
		return nil, err
	}
	erc721Contract, err := contracts.NewERC721Filterer(common.Address{}, nil)
	if err != nil {
		return nil, err
	}

	return &worker{
		ethClient: ethClient,
		lg:        log.GetLogger(),
		mkpAddr:   mkpAddr,

		erc721Abi:      erc721Abi,
		erc721Contract: erc721Contract,
		mkpAbi:         mkpAbi,
		mkpContract:    mkpContract,

		Service: service,
	}, nil
}

func (w *worker) Run(ctx context.Context) error {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go w.listenMkpEvent(ctx, &wg)

	wg.Add(1)
	go w.watchTokenEvent(ctx, &wg)

	wg.Add(1)
	go w.listenOrderExpired(ctx, &wg)

	wg.Wait()
	return nil
}
