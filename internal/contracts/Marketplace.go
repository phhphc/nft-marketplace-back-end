// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TraderListingItem is an auto generated low-level Go binding around an user-defined struct.
type TraderListingItem struct {
	Collection common.Address
	TokenId    *big.Int
	Quantity   *big.Int
	Price      *big.Int
	Seller     common.Address
}

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ListingCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ListingSale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"NewListing\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"cancelListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"getListing\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"internalType\":\"structTrader.ListingItem\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMetaData.ABI instead.
var MarketplaceABI = MarketplaceMetaData.ABI

// Marketplace is an auto generated Go binding around an Ethereum contract.
type Marketplace struct {
	MarketplaceCaller     // Read-only binding to the contract
	MarketplaceTransactor // Write-only binding to the contract
	MarketplaceFilterer   // Log filterer for contract events
}

// MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceSession struct {
	Contract     *Marketplace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceCallerSession struct {
	Contract *MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceTransactorSession struct {
	Contract     *MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceRaw struct {
	Contract *Marketplace // Generic contract binding to access the raw methods on
}

// MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceCallerRaw struct {
	Contract *MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceTransactorRaw struct {
	Contract *MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace creates a new instance of Marketplace, bound to a specific deployed contract.
func NewMarketplace(address common.Address, backend bind.ContractBackend) (*Marketplace, error) {
	contract, err := bindMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// NewMarketplaceCaller creates a new read-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceCaller, error) {
	contract, err := bindMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCaller{contract: contract}, nil
}

// NewMarketplaceTransactor creates a new write-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceTransactor, error) {
	contract, err := bindMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTransactor{contract: contract}, nil
}

// NewMarketplaceFilterer creates a new log filterer instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceFilterer, error) {
	contract, err := bindMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFilterer{contract: contract}, nil
}

// bindMarketplace binds a generic wrapper to an already deployed contract.
func bindMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketplaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transact(opts, method, params...)
}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns((address,uint256,uint256,uint256,address))
func (_Marketplace *MarketplaceCaller) GetListing(opts *bind.CallOpts, listingId *big.Int) (TraderListingItem, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getListing", listingId)

	if err != nil {
		return *new(TraderListingItem), err
	}

	out0 := *abi.ConvertType(out[0], new(TraderListingItem)).(*TraderListingItem)

	return out0, err

}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns((address,uint256,uint256,uint256,address))
func (_Marketplace *MarketplaceSession) GetListing(listingId *big.Int) (TraderListingItem, error) {
	return _Marketplace.Contract.GetListing(&_Marketplace.CallOpts, listingId)
}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns((address,uint256,uint256,uint256,address))
func (_Marketplace *MarketplaceCallerSession) GetListing(listingId *big.Int) (TraderListingItem, error) {
	return _Marketplace.Contract.GetListing(&_Marketplace.CallOpts, listingId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 listingId) payable returns()
func (_Marketplace *MarketplaceTransactor) Buy(opts *bind.TransactOpts, listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "buy", listingId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 listingId) payable returns()
func (_Marketplace *MarketplaceSession) Buy(listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.Buy(&_Marketplace.TransactOpts, listingId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 listingId) payable returns()
func (_Marketplace *MarketplaceTransactorSession) Buy(listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.Buy(&_Marketplace.TransactOpts, listingId)
}

// CancelListing is a paid mutator transaction binding the contract method 0x305a67a8.
//
// Solidity: function cancelListing(uint256 listingId) returns()
func (_Marketplace *MarketplaceTransactor) CancelListing(opts *bind.TransactOpts, listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "cancelListing", listingId)
}

// CancelListing is a paid mutator transaction binding the contract method 0x305a67a8.
//
// Solidity: function cancelListing(uint256 listingId) returns()
func (_Marketplace *MarketplaceSession) CancelListing(listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelListing(&_Marketplace.TransactOpts, listingId)
}

// CancelListing is a paid mutator transaction binding the contract method 0x305a67a8.
//
// Solidity: function cancelListing(uint256 listingId) returns()
func (_Marketplace *MarketplaceTransactorSession) CancelListing(listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelListing(&_Marketplace.TransactOpts, listingId)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Marketplace *MarketplaceTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "onERC721Received", arg0, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Marketplace *MarketplaceSession) OnERC721Received(arg0 common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.OnERC721Received(&_Marketplace.TransactOpts, arg0, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Marketplace *MarketplaceTransactorSession) OnERC721Received(arg0 common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.OnERC721Received(&_Marketplace.TransactOpts, arg0, from, tokenId, data)
}

// MarketplaceListingCanceledIterator is returned from FilterListingCanceled and is used to iterate over the raw logs and unpacked data for ListingCanceled events raised by the Marketplace contract.
type MarketplaceListingCanceledIterator struct {
	Event *MarketplaceListingCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceListingCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceListingCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceListingCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceListingCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceListingCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceListingCanceled represents a ListingCanceled event raised by the Marketplace contract.
type MarketplaceListingCanceled struct {
	ListingId  *big.Int
	Collection common.Address
	TokenId    *big.Int
	Seller     common.Address
	Price      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterListingCanceled is a free log retrieval operation binding the contract event 0xaa2fa98a651c2f4f7c7aa47829192401acec2b7a1596e54871d9f5acfd93f207.
//
// Solidity: event ListingCanceled(uint256 listingId, address indexed collection, uint256 indexed tokenId, address indexed seller, uint256 price)
func (_Marketplace *MarketplaceFilterer) FilterListingCanceled(opts *bind.FilterOpts, collection []common.Address, tokenId []*big.Int, seller []common.Address) (*MarketplaceListingCanceledIterator, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "ListingCanceled", collectionRule, tokenIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceListingCanceledIterator{contract: _Marketplace.contract, event: "ListingCanceled", logs: logs, sub: sub}, nil
}

// WatchListingCanceled is a free log subscription operation binding the contract event 0xaa2fa98a651c2f4f7c7aa47829192401acec2b7a1596e54871d9f5acfd93f207.
//
// Solidity: event ListingCanceled(uint256 listingId, address indexed collection, uint256 indexed tokenId, address indexed seller, uint256 price)
func (_Marketplace *MarketplaceFilterer) WatchListingCanceled(opts *bind.WatchOpts, sink chan<- *MarketplaceListingCanceled, collection []common.Address, tokenId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "ListingCanceled", collectionRule, tokenIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceListingCanceled)
				if err := _Marketplace.contract.UnpackLog(event, "ListingCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseListingCanceled is a log parse operation binding the contract event 0xaa2fa98a651c2f4f7c7aa47829192401acec2b7a1596e54871d9f5acfd93f207.
//
// Solidity: event ListingCanceled(uint256 listingId, address indexed collection, uint256 indexed tokenId, address indexed seller, uint256 price)
func (_Marketplace *MarketplaceFilterer) ParseListingCanceled(log types.Log) (*MarketplaceListingCanceled, error) {
	event := new(MarketplaceListingCanceled)
	if err := _Marketplace.contract.UnpackLog(event, "ListingCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceListingSaleIterator is returned from FilterListingSale and is used to iterate over the raw logs and unpacked data for ListingSale events raised by the Marketplace contract.
type MarketplaceListingSaleIterator struct {
	Event *MarketplaceListingSale // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceListingSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceListingSale)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceListingSale)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceListingSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceListingSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceListingSale represents a ListingSale event raised by the Marketplace contract.
type MarketplaceListingSale struct {
	ListingId  *big.Int
	Collection common.Address
	TokenId    *big.Int
	From       common.Address
	To         common.Address
	Price      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterListingSale is a free log retrieval operation binding the contract event 0x2e074abc18d4bb5b83d782fe1b9d156107564992bf9a5c0292fd937360999620.
//
// Solidity: event ListingSale(uint256 listingId, address indexed collection, uint256 indexed tokenId, address from, address indexed to, uint256 price)
func (_Marketplace *MarketplaceFilterer) FilterListingSale(opts *bind.FilterOpts, collection []common.Address, tokenId []*big.Int, to []common.Address) (*MarketplaceListingSaleIterator, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "ListingSale", collectionRule, tokenIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceListingSaleIterator{contract: _Marketplace.contract, event: "ListingSale", logs: logs, sub: sub}, nil
}

// WatchListingSale is a free log subscription operation binding the contract event 0x2e074abc18d4bb5b83d782fe1b9d156107564992bf9a5c0292fd937360999620.
//
// Solidity: event ListingSale(uint256 listingId, address indexed collection, uint256 indexed tokenId, address from, address indexed to, uint256 price)
func (_Marketplace *MarketplaceFilterer) WatchListingSale(opts *bind.WatchOpts, sink chan<- *MarketplaceListingSale, collection []common.Address, tokenId []*big.Int, to []common.Address) (event.Subscription, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "ListingSale", collectionRule, tokenIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceListingSale)
				if err := _Marketplace.contract.UnpackLog(event, "ListingSale", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseListingSale is a log parse operation binding the contract event 0x2e074abc18d4bb5b83d782fe1b9d156107564992bf9a5c0292fd937360999620.
//
// Solidity: event ListingSale(uint256 listingId, address indexed collection, uint256 indexed tokenId, address from, address indexed to, uint256 price)
func (_Marketplace *MarketplaceFilterer) ParseListingSale(log types.Log) (*MarketplaceListingSale, error) {
	event := new(MarketplaceListingSale)
	if err := _Marketplace.contract.UnpackLog(event, "ListingSale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNewListingIterator is returned from FilterNewListing and is used to iterate over the raw logs and unpacked data for NewListing events raised by the Marketplace contract.
type MarketplaceNewListingIterator struct {
	Event *MarketplaceNewListing // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceNewListingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNewListing)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceNewListing)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceNewListingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNewListingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNewListing represents a NewListing event raised by the Marketplace contract.
type MarketplaceNewListing struct {
	ListingId  *big.Int
	Collection common.Address
	TokenId    *big.Int
	Seller     common.Address
	Price      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewListing is a free log retrieval operation binding the contract event 0xfdafb8fa8c155e0358dae0e8eb9f0494d90a3c9c4da9e209663860ee0c08b8a4.
//
// Solidity: event NewListing(uint256 listingId, address indexed collection, uint256 indexed tokenId, address indexed seller, uint256 price)
func (_Marketplace *MarketplaceFilterer) FilterNewListing(opts *bind.FilterOpts, collection []common.Address, tokenId []*big.Int, seller []common.Address) (*MarketplaceNewListingIterator, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NewListing", collectionRule, tokenIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceNewListingIterator{contract: _Marketplace.contract, event: "NewListing", logs: logs, sub: sub}, nil
}

// WatchNewListing is a free log subscription operation binding the contract event 0xfdafb8fa8c155e0358dae0e8eb9f0494d90a3c9c4da9e209663860ee0c08b8a4.
//
// Solidity: event NewListing(uint256 listingId, address indexed collection, uint256 indexed tokenId, address indexed seller, uint256 price)
func (_Marketplace *MarketplaceFilterer) WatchNewListing(opts *bind.WatchOpts, sink chan<- *MarketplaceNewListing, collection []common.Address, tokenId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NewListing", collectionRule, tokenIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNewListing)
				if err := _Marketplace.contract.UnpackLog(event, "NewListing", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewListing is a log parse operation binding the contract event 0xfdafb8fa8c155e0358dae0e8eb9f0494d90a3c9c4da9e209663860ee0c08b8a4.
//
// Solidity: event NewListing(uint256 listingId, address indexed collection, uint256 indexed tokenId, address indexed seller, uint256 price)
func (_Marketplace *MarketplaceFilterer) ParseNewListing(log types.Log) (*MarketplaceNewListing, error) {
	event := new(MarketplaceNewListing)
	if err := _Marketplace.contract.UnpackLog(event, "NewListing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
