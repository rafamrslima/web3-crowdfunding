// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crowdfunding

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
	_ = abi.ConvertType
)

// CrowdFundingCampaign is an auto generated low-level Go binding around an user-defined struct.
type CrowdFundingCampaign struct {
	Owner           common.Address
	Title           string
	Description     string
	Target          *big.Int
	Deadline        *big.Int
	AmountCollected *big.Int
	Image           string
	Donators        []common.Address
	Donations       []*big.Int
	Withdrawn       bool
}

// CrowdfundingMetaData contains all meta data concerning the Crowdfunding contract.
var CrowdfundingMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"campaigns\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"title\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"target\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountCollected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"image\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"withdrawn\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createCampaign\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_title\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_description\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_target\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_image\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"donateToCampaign\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getCampaigns\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCrowdFunding.Campaign[]\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"title\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"target\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountCollected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"image\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"donators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"donations\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"withdrawn\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDonators\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numberOfCampaigns\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"refundDonators\",\"inputs\":[{\"name\":\"_idCampaign\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"_idCampaign\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CampaignCreated\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"title\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"targetWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CampaignWithdrawn\",\"inputs\":[{\"name\":\"campaignId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DonationReceived\",\"inputs\":[{\"name\":\"campaignId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// CrowdfundingABI is the input ABI used to generate the binding from.
// Deprecated: Use CrowdfundingMetaData.ABI instead.
var CrowdfundingABI = CrowdfundingMetaData.ABI

// Crowdfunding is an auto generated Go binding around an Ethereum contract.
type Crowdfunding struct {
	CrowdfundingCaller     // Read-only binding to the contract
	CrowdfundingTransactor // Write-only binding to the contract
	CrowdfundingFilterer   // Log filterer for contract events
}

// CrowdfundingCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrowdfundingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdfundingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrowdfundingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdfundingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrowdfundingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdfundingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrowdfundingSession struct {
	Contract     *Crowdfunding     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrowdfundingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrowdfundingCallerSession struct {
	Contract *CrowdfundingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CrowdfundingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrowdfundingTransactorSession struct {
	Contract     *CrowdfundingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CrowdfundingRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrowdfundingRaw struct {
	Contract *Crowdfunding // Generic contract binding to access the raw methods on
}

// CrowdfundingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrowdfundingCallerRaw struct {
	Contract *CrowdfundingCaller // Generic read-only contract binding to access the raw methods on
}

// CrowdfundingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrowdfundingTransactorRaw struct {
	Contract *CrowdfundingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrowdfunding creates a new instance of Crowdfunding, bound to a specific deployed contract.
func NewCrowdfunding(address common.Address, backend bind.ContractBackend) (*Crowdfunding, error) {
	contract, err := bindCrowdfunding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Crowdfunding{CrowdfundingCaller: CrowdfundingCaller{contract: contract}, CrowdfundingTransactor: CrowdfundingTransactor{contract: contract}, CrowdfundingFilterer: CrowdfundingFilterer{contract: contract}}, nil
}

// NewCrowdfundingCaller creates a new read-only instance of Crowdfunding, bound to a specific deployed contract.
func NewCrowdfundingCaller(address common.Address, caller bind.ContractCaller) (*CrowdfundingCaller, error) {
	contract, err := bindCrowdfunding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrowdfundingCaller{contract: contract}, nil
}

// NewCrowdfundingTransactor creates a new write-only instance of Crowdfunding, bound to a specific deployed contract.
func NewCrowdfundingTransactor(address common.Address, transactor bind.ContractTransactor) (*CrowdfundingTransactor, error) {
	contract, err := bindCrowdfunding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrowdfundingTransactor{contract: contract}, nil
}

// NewCrowdfundingFilterer creates a new log filterer instance of Crowdfunding, bound to a specific deployed contract.
func NewCrowdfundingFilterer(address common.Address, filterer bind.ContractFilterer) (*CrowdfundingFilterer, error) {
	contract, err := bindCrowdfunding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrowdfundingFilterer{contract: contract}, nil
}

// bindCrowdfunding binds a generic wrapper to an already deployed contract.
func bindCrowdfunding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrowdfundingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crowdfunding *CrowdfundingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crowdfunding.Contract.CrowdfundingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crowdfunding *CrowdfundingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crowdfunding.Contract.CrowdfundingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crowdfunding *CrowdfundingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crowdfunding.Contract.CrowdfundingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crowdfunding *CrowdfundingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crowdfunding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crowdfunding *CrowdfundingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crowdfunding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crowdfunding *CrowdfundingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crowdfunding.Contract.contract.Transact(opts, method, params...)
}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns(uint256 ) view returns(address owner, string title, string description, uint256 target, uint256 deadline, uint256 amountCollected, string image, bool withdrawn)
func (_Crowdfunding *CrowdfundingCaller) Campaigns(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner           common.Address
	Title           string
	Description     string
	Target          *big.Int
	Deadline        *big.Int
	AmountCollected *big.Int
	Image           string
	Withdrawn       bool
}, error) {
	var out []interface{}
	err := _Crowdfunding.contract.Call(opts, &out, "campaigns", arg0)

	outstruct := new(struct {
		Owner           common.Address
		Title           string
		Description     string
		Target          *big.Int
		Deadline        *big.Int
		AmountCollected *big.Int
		Image           string
		Withdrawn       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Title = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Target = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Deadline = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AmountCollected = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Image = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.Withdrawn = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns(uint256 ) view returns(address owner, string title, string description, uint256 target, uint256 deadline, uint256 amountCollected, string image, bool withdrawn)
func (_Crowdfunding *CrowdfundingSession) Campaigns(arg0 *big.Int) (struct {
	Owner           common.Address
	Title           string
	Description     string
	Target          *big.Int
	Deadline        *big.Int
	AmountCollected *big.Int
	Image           string
	Withdrawn       bool
}, error) {
	return _Crowdfunding.Contract.Campaigns(&_Crowdfunding.CallOpts, arg0)
}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns(uint256 ) view returns(address owner, string title, string description, uint256 target, uint256 deadline, uint256 amountCollected, string image, bool withdrawn)
func (_Crowdfunding *CrowdfundingCallerSession) Campaigns(arg0 *big.Int) (struct {
	Owner           common.Address
	Title           string
	Description     string
	Target          *big.Int
	Deadline        *big.Int
	AmountCollected *big.Int
	Image           string
	Withdrawn       bool
}, error) {
	return _Crowdfunding.Contract.Campaigns(&_Crowdfunding.CallOpts, arg0)
}

// GetCampaigns is a free data retrieval call binding the contract method 0xa6b03633.
//
// Solidity: function getCampaigns() view returns((address,string,string,uint256,uint256,uint256,string,address[],uint256[],bool)[])
func (_Crowdfunding *CrowdfundingCaller) GetCampaigns(opts *bind.CallOpts) ([]CrowdFundingCampaign, error) {
	var out []interface{}
	err := _Crowdfunding.contract.Call(opts, &out, "getCampaigns")

	if err != nil {
		return *new([]CrowdFundingCampaign), err
	}

	out0 := *abi.ConvertType(out[0], new([]CrowdFundingCampaign)).(*[]CrowdFundingCampaign)

	return out0, err

}

// GetCampaigns is a free data retrieval call binding the contract method 0xa6b03633.
//
// Solidity: function getCampaigns() view returns((address,string,string,uint256,uint256,uint256,string,address[],uint256[],bool)[])
func (_Crowdfunding *CrowdfundingSession) GetCampaigns() ([]CrowdFundingCampaign, error) {
	return _Crowdfunding.Contract.GetCampaigns(&_Crowdfunding.CallOpts)
}

// GetCampaigns is a free data retrieval call binding the contract method 0xa6b03633.
//
// Solidity: function getCampaigns() view returns((address,string,string,uint256,uint256,uint256,string,address[],uint256[],bool)[])
func (_Crowdfunding *CrowdfundingCallerSession) GetCampaigns() ([]CrowdFundingCampaign, error) {
	return _Crowdfunding.Contract.GetCampaigns(&_Crowdfunding.CallOpts)
}

// GetDonators is a free data retrieval call binding the contract method 0x0fa91fa9.
//
// Solidity: function getDonators(uint256 _id) view returns(address[], uint256[])
func (_Crowdfunding *CrowdfundingCaller) GetDonators(opts *bind.CallOpts, _id *big.Int) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Crowdfunding.contract.Call(opts, &out, "getDonators", _id)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetDonators is a free data retrieval call binding the contract method 0x0fa91fa9.
//
// Solidity: function getDonators(uint256 _id) view returns(address[], uint256[])
func (_Crowdfunding *CrowdfundingSession) GetDonators(_id *big.Int) ([]common.Address, []*big.Int, error) {
	return _Crowdfunding.Contract.GetDonators(&_Crowdfunding.CallOpts, _id)
}

// GetDonators is a free data retrieval call binding the contract method 0x0fa91fa9.
//
// Solidity: function getDonators(uint256 _id) view returns(address[], uint256[])
func (_Crowdfunding *CrowdfundingCallerSession) GetDonators(_id *big.Int) ([]common.Address, []*big.Int, error) {
	return _Crowdfunding.Contract.GetDonators(&_Crowdfunding.CallOpts, _id)
}

// NumberOfCampaigns is a free data retrieval call binding the contract method 0x07ca140d.
//
// Solidity: function numberOfCampaigns() view returns(uint256)
func (_Crowdfunding *CrowdfundingCaller) NumberOfCampaigns(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crowdfunding.contract.Call(opts, &out, "numberOfCampaigns")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfCampaigns is a free data retrieval call binding the contract method 0x07ca140d.
//
// Solidity: function numberOfCampaigns() view returns(uint256)
func (_Crowdfunding *CrowdfundingSession) NumberOfCampaigns() (*big.Int, error) {
	return _Crowdfunding.Contract.NumberOfCampaigns(&_Crowdfunding.CallOpts)
}

// NumberOfCampaigns is a free data retrieval call binding the contract method 0x07ca140d.
//
// Solidity: function numberOfCampaigns() view returns(uint256)
func (_Crowdfunding *CrowdfundingCallerSession) NumberOfCampaigns() (*big.Int, error) {
	return _Crowdfunding.Contract.NumberOfCampaigns(&_Crowdfunding.CallOpts)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0x9943e3a1.
//
// Solidity: function createCampaign(address _owner, string _title, string _description, uint256 _target, uint256 _deadline, string _image) returns(uint256)
func (_Crowdfunding *CrowdfundingTransactor) CreateCampaign(opts *bind.TransactOpts, _owner common.Address, _title string, _description string, _target *big.Int, _deadline *big.Int, _image string) (*types.Transaction, error) {
	return _Crowdfunding.contract.Transact(opts, "createCampaign", _owner, _title, _description, _target, _deadline, _image)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0x9943e3a1.
//
// Solidity: function createCampaign(address _owner, string _title, string _description, uint256 _target, uint256 _deadline, string _image) returns(uint256)
func (_Crowdfunding *CrowdfundingSession) CreateCampaign(_owner common.Address, _title string, _description string, _target *big.Int, _deadline *big.Int, _image string) (*types.Transaction, error) {
	return _Crowdfunding.Contract.CreateCampaign(&_Crowdfunding.TransactOpts, _owner, _title, _description, _target, _deadline, _image)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0x9943e3a1.
//
// Solidity: function createCampaign(address _owner, string _title, string _description, uint256 _target, uint256 _deadline, string _image) returns(uint256)
func (_Crowdfunding *CrowdfundingTransactorSession) CreateCampaign(_owner common.Address, _title string, _description string, _target *big.Int, _deadline *big.Int, _image string) (*types.Transaction, error) {
	return _Crowdfunding.Contract.CreateCampaign(&_Crowdfunding.TransactOpts, _owner, _title, _description, _target, _deadline, _image)
}

// DonateToCampaign is a paid mutator transaction binding the contract method 0x42a4fda8.
//
// Solidity: function donateToCampaign(uint256 _id) payable returns()
func (_Crowdfunding *CrowdfundingTransactor) DonateToCampaign(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Crowdfunding.contract.Transact(opts, "donateToCampaign", _id)
}

// DonateToCampaign is a paid mutator transaction binding the contract method 0x42a4fda8.
//
// Solidity: function donateToCampaign(uint256 _id) payable returns()
func (_Crowdfunding *CrowdfundingSession) DonateToCampaign(_id *big.Int) (*types.Transaction, error) {
	return _Crowdfunding.Contract.DonateToCampaign(&_Crowdfunding.TransactOpts, _id)
}

// DonateToCampaign is a paid mutator transaction binding the contract method 0x42a4fda8.
//
// Solidity: function donateToCampaign(uint256 _id) payable returns()
func (_Crowdfunding *CrowdfundingTransactorSession) DonateToCampaign(_id *big.Int) (*types.Transaction, error) {
	return _Crowdfunding.Contract.DonateToCampaign(&_Crowdfunding.TransactOpts, _id)
}

// RefundDonators is a paid mutator transaction binding the contract method 0xd4a031b1.
//
// Solidity: function refundDonators(uint256 _idCampaign) payable returns()
func (_Crowdfunding *CrowdfundingTransactor) RefundDonators(opts *bind.TransactOpts, _idCampaign *big.Int) (*types.Transaction, error) {
	return _Crowdfunding.contract.Transact(opts, "refundDonators", _idCampaign)
}

// RefundDonators is a paid mutator transaction binding the contract method 0xd4a031b1.
//
// Solidity: function refundDonators(uint256 _idCampaign) payable returns()
func (_Crowdfunding *CrowdfundingSession) RefundDonators(_idCampaign *big.Int) (*types.Transaction, error) {
	return _Crowdfunding.Contract.RefundDonators(&_Crowdfunding.TransactOpts, _idCampaign)
}

// RefundDonators is a paid mutator transaction binding the contract method 0xd4a031b1.
//
// Solidity: function refundDonators(uint256 _idCampaign) payable returns()
func (_Crowdfunding *CrowdfundingTransactorSession) RefundDonators(_idCampaign *big.Int) (*types.Transaction, error) {
	return _Crowdfunding.Contract.RefundDonators(&_Crowdfunding.TransactOpts, _idCampaign)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _idCampaign) returns()
func (_Crowdfunding *CrowdfundingTransactor) Withdraw(opts *bind.TransactOpts, _idCampaign *big.Int) (*types.Transaction, error) {
	return _Crowdfunding.contract.Transact(opts, "withdraw", _idCampaign)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _idCampaign) returns()
func (_Crowdfunding *CrowdfundingSession) Withdraw(_idCampaign *big.Int) (*types.Transaction, error) {
	return _Crowdfunding.Contract.Withdraw(&_Crowdfunding.TransactOpts, _idCampaign)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _idCampaign) returns()
func (_Crowdfunding *CrowdfundingTransactorSession) Withdraw(_idCampaign *big.Int) (*types.Transaction, error) {
	return _Crowdfunding.Contract.Withdraw(&_Crowdfunding.TransactOpts, _idCampaign)
}

// CrowdfundingCampaignCreatedIterator is returned from FilterCampaignCreated and is used to iterate over the raw logs and unpacked data for CampaignCreated events raised by the Crowdfunding contract.
type CrowdfundingCampaignCreatedIterator struct {
	Event *CrowdfundingCampaignCreated // Event containing the contract specifics and raw log

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
func (it *CrowdfundingCampaignCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdfundingCampaignCreated)
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
		it.Event = new(CrowdfundingCampaignCreated)
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
func (it *CrowdfundingCampaignCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdfundingCampaignCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdfundingCampaignCreated represents a CampaignCreated event raised by the Crowdfunding contract.
type CrowdfundingCampaignCreated struct {
	Id        *big.Int
	Owner     common.Address
	Title     string
	TargetWei *big.Int
	Deadline  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCampaignCreated is a free log retrieval operation binding the contract event 0xdc26653af5b99b2da33e2ad69ee6600d9aeccc82b034501db4338309615ca238.
//
// Solidity: event CampaignCreated(uint256 indexed id, address indexed owner, string title, uint256 targetWei, uint256 deadline)
func (_Crowdfunding *CrowdfundingFilterer) FilterCampaignCreated(opts *bind.FilterOpts, id []*big.Int, owner []common.Address) (*CrowdfundingCampaignCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Crowdfunding.contract.FilterLogs(opts, "CampaignCreated", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CrowdfundingCampaignCreatedIterator{contract: _Crowdfunding.contract, event: "CampaignCreated", logs: logs, sub: sub}, nil
}

// WatchCampaignCreated is a free log subscription operation binding the contract event 0xdc26653af5b99b2da33e2ad69ee6600d9aeccc82b034501db4338309615ca238.
//
// Solidity: event CampaignCreated(uint256 indexed id, address indexed owner, string title, uint256 targetWei, uint256 deadline)
func (_Crowdfunding *CrowdfundingFilterer) WatchCampaignCreated(opts *bind.WatchOpts, sink chan<- *CrowdfundingCampaignCreated, id []*big.Int, owner []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Crowdfunding.contract.WatchLogs(opts, "CampaignCreated", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdfundingCampaignCreated)
				if err := _Crowdfunding.contract.UnpackLog(event, "CampaignCreated", log); err != nil {
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

// ParseCampaignCreated is a log parse operation binding the contract event 0xdc26653af5b99b2da33e2ad69ee6600d9aeccc82b034501db4338309615ca238.
//
// Solidity: event CampaignCreated(uint256 indexed id, address indexed owner, string title, uint256 targetWei, uint256 deadline)
func (_Crowdfunding *CrowdfundingFilterer) ParseCampaignCreated(log types.Log) (*CrowdfundingCampaignCreated, error) {
	event := new(CrowdfundingCampaignCreated)
	if err := _Crowdfunding.contract.UnpackLog(event, "CampaignCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdfundingCampaignWithdrawnIterator is returned from FilterCampaignWithdrawn and is used to iterate over the raw logs and unpacked data for CampaignWithdrawn events raised by the Crowdfunding contract.
type CrowdfundingCampaignWithdrawnIterator struct {
	Event *CrowdfundingCampaignWithdrawn // Event containing the contract specifics and raw log

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
func (it *CrowdfundingCampaignWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdfundingCampaignWithdrawn)
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
		it.Event = new(CrowdfundingCampaignWithdrawn)
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
func (it *CrowdfundingCampaignWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdfundingCampaignWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdfundingCampaignWithdrawn represents a CampaignWithdrawn event raised by the Crowdfunding contract.
type CrowdfundingCampaignWithdrawn struct {
	CampaignId *big.Int
	Owner      common.Address
	AmountWei  *big.Int
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCampaignWithdrawn is a free log retrieval operation binding the contract event 0xf6855da62af7f8886a08b21d9396963710107b9819475543c2ab579138620bf8.
//
// Solidity: event CampaignWithdrawn(uint256 indexed campaignId, address indexed owner, uint256 amountWei, uint256 timestamp)
func (_Crowdfunding *CrowdfundingFilterer) FilterCampaignWithdrawn(opts *bind.FilterOpts, campaignId []*big.Int, owner []common.Address) (*CrowdfundingCampaignWithdrawnIterator, error) {

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Crowdfunding.contract.FilterLogs(opts, "CampaignWithdrawn", campaignIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CrowdfundingCampaignWithdrawnIterator{contract: _Crowdfunding.contract, event: "CampaignWithdrawn", logs: logs, sub: sub}, nil
}

// WatchCampaignWithdrawn is a free log subscription operation binding the contract event 0xf6855da62af7f8886a08b21d9396963710107b9819475543c2ab579138620bf8.
//
// Solidity: event CampaignWithdrawn(uint256 indexed campaignId, address indexed owner, uint256 amountWei, uint256 timestamp)
func (_Crowdfunding *CrowdfundingFilterer) WatchCampaignWithdrawn(opts *bind.WatchOpts, sink chan<- *CrowdfundingCampaignWithdrawn, campaignId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Crowdfunding.contract.WatchLogs(opts, "CampaignWithdrawn", campaignIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdfundingCampaignWithdrawn)
				if err := _Crowdfunding.contract.UnpackLog(event, "CampaignWithdrawn", log); err != nil {
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

// ParseCampaignWithdrawn is a log parse operation binding the contract event 0xf6855da62af7f8886a08b21d9396963710107b9819475543c2ab579138620bf8.
//
// Solidity: event CampaignWithdrawn(uint256 indexed campaignId, address indexed owner, uint256 amountWei, uint256 timestamp)
func (_Crowdfunding *CrowdfundingFilterer) ParseCampaignWithdrawn(log types.Log) (*CrowdfundingCampaignWithdrawn, error) {
	event := new(CrowdfundingCampaignWithdrawn)
	if err := _Crowdfunding.contract.UnpackLog(event, "CampaignWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdfundingDonationReceivedIterator is returned from FilterDonationReceived and is used to iterate over the raw logs and unpacked data for DonationReceived events raised by the Crowdfunding contract.
type CrowdfundingDonationReceivedIterator struct {
	Event *CrowdfundingDonationReceived // Event containing the contract specifics and raw log

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
func (it *CrowdfundingDonationReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdfundingDonationReceived)
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
		it.Event = new(CrowdfundingDonationReceived)
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
func (it *CrowdfundingDonationReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdfundingDonationReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdfundingDonationReceived represents a DonationReceived event raised by the Crowdfunding contract.
type CrowdfundingDonationReceived struct {
	CampaignId *big.Int
	Receiver   common.Address
	Donor      common.Address
	AmountWei  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDonationReceived is a free log retrieval operation binding the contract event 0x43287c3057ca71c2f01d2236b70a9e5376f802564d6de4e42634af8cf0e3f18a.
//
// Solidity: event DonationReceived(uint256 indexed campaignId, address indexed receiver, address indexed donor, uint256 amountWei)
func (_Crowdfunding *CrowdfundingFilterer) FilterDonationReceived(opts *bind.FilterOpts, campaignId []*big.Int, receiver []common.Address, donor []common.Address) (*CrowdfundingDonationReceivedIterator, error) {

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _Crowdfunding.contract.FilterLogs(opts, "DonationReceived", campaignIdRule, receiverRule, donorRule)
	if err != nil {
		return nil, err
	}
	return &CrowdfundingDonationReceivedIterator{contract: _Crowdfunding.contract, event: "DonationReceived", logs: logs, sub: sub}, nil
}

// WatchDonationReceived is a free log subscription operation binding the contract event 0x43287c3057ca71c2f01d2236b70a9e5376f802564d6de4e42634af8cf0e3f18a.
//
// Solidity: event DonationReceived(uint256 indexed campaignId, address indexed receiver, address indexed donor, uint256 amountWei)
func (_Crowdfunding *CrowdfundingFilterer) WatchDonationReceived(opts *bind.WatchOpts, sink chan<- *CrowdfundingDonationReceived, campaignId []*big.Int, receiver []common.Address, donor []common.Address) (event.Subscription, error) {

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _Crowdfunding.contract.WatchLogs(opts, "DonationReceived", campaignIdRule, receiverRule, donorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdfundingDonationReceived)
				if err := _Crowdfunding.contract.UnpackLog(event, "DonationReceived", log); err != nil {
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

// ParseDonationReceived is a log parse operation binding the contract event 0x43287c3057ca71c2f01d2236b70a9e5376f802564d6de4e42634af8cf0e3f18a.
//
// Solidity: event DonationReceived(uint256 indexed campaignId, address indexed receiver, address indexed donor, uint256 amountWei)
func (_Crowdfunding *CrowdfundingFilterer) ParseDonationReceived(log types.Log) (*CrowdfundingDonationReceived, error) {
	event := new(CrowdfundingDonationReceived)
	if err := _Crowdfunding.contract.UnpackLog(event, "DonationReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
