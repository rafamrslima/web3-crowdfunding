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
	Target          *big.Int
	Deadline        *big.Int
	AmountCollected *big.Int
	Donators        []common.Address
	Donations       []*big.Int
	Withdrawn       bool
}

// CrowdfundingMetaData contains all meta data concerning the Crowdfunding contract.
var CrowdfundingMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_usdc\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"campaigns\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountCollected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"withdrawn\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contributions\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createCampaign\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_target\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"donateToCampaign\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCampaigns\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCrowdFunding.Campaign[]\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountCollected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"donators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"donations\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"withdrawn\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDonators\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numberOfCampaigns\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"refundDonor\",\"inputs\":[{\"name\":\"_idCampaign\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"usdc\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"_idCampaign\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CampaignCreated\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DonationReceived\",\"inputs\":[{\"name\":\"campaignId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"donor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DonationRefunded\",\"inputs\":[{\"name\":\"campaignId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"donor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"totalContributed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundsWithdrawn\",\"inputs\":[{\"name\":\"campaignId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
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
// Solidity: function campaigns(uint256 ) view returns(address owner, uint256 target, uint256 deadline, uint256 amountCollected, bool withdrawn)
func (_Crowdfunding *CrowdfundingCaller) Campaigns(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner           common.Address
	Target          *big.Int
	Deadline        *big.Int
	AmountCollected *big.Int
	Withdrawn       bool
}, error) {
	var out []interface{}
	err := _Crowdfunding.contract.Call(opts, &out, "campaigns", arg0)

	outstruct := new(struct {
		Owner           common.Address
		Target          *big.Int
		Deadline        *big.Int
		AmountCollected *big.Int
		Withdrawn       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Target = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Deadline = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AmountCollected = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Withdrawn = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns(uint256 ) view returns(address owner, uint256 target, uint256 deadline, uint256 amountCollected, bool withdrawn)
func (_Crowdfunding *CrowdfundingSession) Campaigns(arg0 *big.Int) (struct {
	Owner           common.Address
	Target          *big.Int
	Deadline        *big.Int
	AmountCollected *big.Int
	Withdrawn       bool
}, error) {
	return _Crowdfunding.Contract.Campaigns(&_Crowdfunding.CallOpts, arg0)
}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns(uint256 ) view returns(address owner, uint256 target, uint256 deadline, uint256 amountCollected, bool withdrawn)
func (_Crowdfunding *CrowdfundingCallerSession) Campaigns(arg0 *big.Int) (struct {
	Owner           common.Address
	Target          *big.Int
	Deadline        *big.Int
	AmountCollected *big.Int
	Withdrawn       bool
}, error) {
	return _Crowdfunding.Contract.Campaigns(&_Crowdfunding.CallOpts, arg0)
}

// Contributions is a free data retrieval call binding the contract method 0x3d891f59.
//
// Solidity: function contributions(uint256 , address ) view returns(uint256)
func (_Crowdfunding *CrowdfundingCaller) Contributions(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Crowdfunding.contract.Call(opts, &out, "contributions", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Contributions is a free data retrieval call binding the contract method 0x3d891f59.
//
// Solidity: function contributions(uint256 , address ) view returns(uint256)
func (_Crowdfunding *CrowdfundingSession) Contributions(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Crowdfunding.Contract.Contributions(&_Crowdfunding.CallOpts, arg0, arg1)
}

// Contributions is a free data retrieval call binding the contract method 0x3d891f59.
//
// Solidity: function contributions(uint256 , address ) view returns(uint256)
func (_Crowdfunding *CrowdfundingCallerSession) Contributions(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Crowdfunding.Contract.Contributions(&_Crowdfunding.CallOpts, arg0, arg1)
}

// GetCampaigns is a free data retrieval call binding the contract method 0xa6b03633.
//
// Solidity: function getCampaigns() view returns((address,uint256,uint256,uint256,address[],uint256[],bool)[])
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
// Solidity: function getCampaigns() view returns((address,uint256,uint256,uint256,address[],uint256[],bool)[])
func (_Crowdfunding *CrowdfundingSession) GetCampaigns() ([]CrowdFundingCampaign, error) {
	return _Crowdfunding.Contract.GetCampaigns(&_Crowdfunding.CallOpts)
}

// GetCampaigns is a free data retrieval call binding the contract method 0xa6b03633.
//
// Solidity: function getCampaigns() view returns((address,uint256,uint256,uint256,address[],uint256[],bool)[])
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

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_Crowdfunding *CrowdfundingCaller) Usdc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crowdfunding.contract.Call(opts, &out, "usdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_Crowdfunding *CrowdfundingSession) Usdc() (common.Address, error) {
	return _Crowdfunding.Contract.Usdc(&_Crowdfunding.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_Crowdfunding *CrowdfundingCallerSession) Usdc() (common.Address, error) {
	return _Crowdfunding.Contract.Usdc(&_Crowdfunding.CallOpts)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0x461b6004.
//
// Solidity: function createCampaign(address _owner, uint256 _target, uint256 _deadline) returns(uint256)
func (_Crowdfunding *CrowdfundingTransactor) CreateCampaign(opts *bind.TransactOpts, _owner common.Address, _target *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _Crowdfunding.contract.Transact(opts, "createCampaign", _owner, _target, _deadline)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0x461b6004.
//
// Solidity: function createCampaign(address _owner, uint256 _target, uint256 _deadline) returns(uint256)
func (_Crowdfunding *CrowdfundingSession) CreateCampaign(_owner common.Address, _target *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _Crowdfunding.Contract.CreateCampaign(&_Crowdfunding.TransactOpts, _owner, _target, _deadline)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0x461b6004.
//
// Solidity: function createCampaign(address _owner, uint256 _target, uint256 _deadline) returns(uint256)
func (_Crowdfunding *CrowdfundingTransactorSession) CreateCampaign(_owner common.Address, _target *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _Crowdfunding.Contract.CreateCampaign(&_Crowdfunding.TransactOpts, _owner, _target, _deadline)
}

// DonateToCampaign is a paid mutator transaction binding the contract method 0x2b7216e5.
//
// Solidity: function donateToCampaign(uint256 _id, uint256 amount) returns()
func (_Crowdfunding *CrowdfundingTransactor) DonateToCampaign(opts *bind.TransactOpts, _id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Crowdfunding.contract.Transact(opts, "donateToCampaign", _id, amount)
}

// DonateToCampaign is a paid mutator transaction binding the contract method 0x2b7216e5.
//
// Solidity: function donateToCampaign(uint256 _id, uint256 amount) returns()
func (_Crowdfunding *CrowdfundingSession) DonateToCampaign(_id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Crowdfunding.Contract.DonateToCampaign(&_Crowdfunding.TransactOpts, _id, amount)
}

// DonateToCampaign is a paid mutator transaction binding the contract method 0x2b7216e5.
//
// Solidity: function donateToCampaign(uint256 _id, uint256 amount) returns()
func (_Crowdfunding *CrowdfundingTransactorSession) DonateToCampaign(_id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Crowdfunding.Contract.DonateToCampaign(&_Crowdfunding.TransactOpts, _id, amount)
}

// RefundDonor is a paid mutator transaction binding the contract method 0x01bc3f6d.
//
// Solidity: function refundDonor(uint256 _idCampaign) returns()
func (_Crowdfunding *CrowdfundingTransactor) RefundDonor(opts *bind.TransactOpts, _idCampaign *big.Int) (*types.Transaction, error) {
	return _Crowdfunding.contract.Transact(opts, "refundDonor", _idCampaign)
}

// RefundDonor is a paid mutator transaction binding the contract method 0x01bc3f6d.
//
// Solidity: function refundDonor(uint256 _idCampaign) returns()
func (_Crowdfunding *CrowdfundingSession) RefundDonor(_idCampaign *big.Int) (*types.Transaction, error) {
	return _Crowdfunding.Contract.RefundDonor(&_Crowdfunding.TransactOpts, _idCampaign)
}

// RefundDonor is a paid mutator transaction binding the contract method 0x01bc3f6d.
//
// Solidity: function refundDonor(uint256 _idCampaign) returns()
func (_Crowdfunding *CrowdfundingTransactorSession) RefundDonor(_idCampaign *big.Int) (*types.Transaction, error) {
	return _Crowdfunding.Contract.RefundDonor(&_Crowdfunding.TransactOpts, _idCampaign)
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
	Id       *big.Int
	Owner    common.Address
	Target   *big.Int
	Deadline *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCampaignCreated is a free log retrieval operation binding the contract event 0x91b289a829e71d811b8c69e4a24ba2d40d115d8a236e9a724cb3bb2d43cf7223.
//
// Solidity: event CampaignCreated(uint256 indexed id, address indexed owner, uint256 target, uint256 deadline)
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

// WatchCampaignCreated is a free log subscription operation binding the contract event 0x91b289a829e71d811b8c69e4a24ba2d40d115d8a236e9a724cb3bb2d43cf7223.
//
// Solidity: event CampaignCreated(uint256 indexed id, address indexed owner, uint256 target, uint256 deadline)
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

// ParseCampaignCreated is a log parse operation binding the contract event 0x91b289a829e71d811b8c69e4a24ba2d40d115d8a236e9a724cb3bb2d43cf7223.
//
// Solidity: event CampaignCreated(uint256 indexed id, address indexed owner, uint256 target, uint256 deadline)
func (_Crowdfunding *CrowdfundingFilterer) ParseCampaignCreated(log types.Log) (*CrowdfundingCampaignCreated, error) {
	event := new(CrowdfundingCampaignCreated)
	if err := _Crowdfunding.contract.UnpackLog(event, "CampaignCreated", log); err != nil {
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
	Donor      common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDonationReceived is a free log retrieval operation binding the contract event 0x0b5b4c52969ff7329ecf7ee536409fda87812b15a8622bc6e8cdeab3aee14a26.
//
// Solidity: event DonationReceived(uint256 indexed campaignId, address indexed donor, uint256 amount)
func (_Crowdfunding *CrowdfundingFilterer) FilterDonationReceived(opts *bind.FilterOpts, campaignId []*big.Int, donor []common.Address) (*CrowdfundingDonationReceivedIterator, error) {

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}
	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _Crowdfunding.contract.FilterLogs(opts, "DonationReceived", campaignIdRule, donorRule)
	if err != nil {
		return nil, err
	}
	return &CrowdfundingDonationReceivedIterator{contract: _Crowdfunding.contract, event: "DonationReceived", logs: logs, sub: sub}, nil
}

// WatchDonationReceived is a free log subscription operation binding the contract event 0x0b5b4c52969ff7329ecf7ee536409fda87812b15a8622bc6e8cdeab3aee14a26.
//
// Solidity: event DonationReceived(uint256 indexed campaignId, address indexed donor, uint256 amount)
func (_Crowdfunding *CrowdfundingFilterer) WatchDonationReceived(opts *bind.WatchOpts, sink chan<- *CrowdfundingDonationReceived, campaignId []*big.Int, donor []common.Address) (event.Subscription, error) {

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}
	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _Crowdfunding.contract.WatchLogs(opts, "DonationReceived", campaignIdRule, donorRule)
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

// ParseDonationReceived is a log parse operation binding the contract event 0x0b5b4c52969ff7329ecf7ee536409fda87812b15a8622bc6e8cdeab3aee14a26.
//
// Solidity: event DonationReceived(uint256 indexed campaignId, address indexed donor, uint256 amount)
func (_Crowdfunding *CrowdfundingFilterer) ParseDonationReceived(log types.Log) (*CrowdfundingDonationReceived, error) {
	event := new(CrowdfundingDonationReceived)
	if err := _Crowdfunding.contract.UnpackLog(event, "DonationReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdfundingDonationRefundedIterator is returned from FilterDonationRefunded and is used to iterate over the raw logs and unpacked data for DonationRefunded events raised by the Crowdfunding contract.
type CrowdfundingDonationRefundedIterator struct {
	Event *CrowdfundingDonationRefunded // Event containing the contract specifics and raw log

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
func (it *CrowdfundingDonationRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdfundingDonationRefunded)
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
		it.Event = new(CrowdfundingDonationRefunded)
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
func (it *CrowdfundingDonationRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdfundingDonationRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdfundingDonationRefunded represents a DonationRefunded event raised by the Crowdfunding contract.
type CrowdfundingDonationRefunded struct {
	CampaignId       *big.Int
	Donor            common.Address
	TotalContributed *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDonationRefunded is a free log retrieval operation binding the contract event 0x4152e14791691379e08bbc1c4beeb7197ecbf5db73dc45a8c3f946f66a931c03.
//
// Solidity: event DonationRefunded(uint256 indexed campaignId, address indexed donor, uint256 totalContributed)
func (_Crowdfunding *CrowdfundingFilterer) FilterDonationRefunded(opts *bind.FilterOpts, campaignId []*big.Int, donor []common.Address) (*CrowdfundingDonationRefundedIterator, error) {

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}
	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _Crowdfunding.contract.FilterLogs(opts, "DonationRefunded", campaignIdRule, donorRule)
	if err != nil {
		return nil, err
	}
	return &CrowdfundingDonationRefundedIterator{contract: _Crowdfunding.contract, event: "DonationRefunded", logs: logs, sub: sub}, nil
}

// WatchDonationRefunded is a free log subscription operation binding the contract event 0x4152e14791691379e08bbc1c4beeb7197ecbf5db73dc45a8c3f946f66a931c03.
//
// Solidity: event DonationRefunded(uint256 indexed campaignId, address indexed donor, uint256 totalContributed)
func (_Crowdfunding *CrowdfundingFilterer) WatchDonationRefunded(opts *bind.WatchOpts, sink chan<- *CrowdfundingDonationRefunded, campaignId []*big.Int, donor []common.Address) (event.Subscription, error) {

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}
	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _Crowdfunding.contract.WatchLogs(opts, "DonationRefunded", campaignIdRule, donorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdfundingDonationRefunded)
				if err := _Crowdfunding.contract.UnpackLog(event, "DonationRefunded", log); err != nil {
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

// ParseDonationRefunded is a log parse operation binding the contract event 0x4152e14791691379e08bbc1c4beeb7197ecbf5db73dc45a8c3f946f66a931c03.
//
// Solidity: event DonationRefunded(uint256 indexed campaignId, address indexed donor, uint256 totalContributed)
func (_Crowdfunding *CrowdfundingFilterer) ParseDonationRefunded(log types.Log) (*CrowdfundingDonationRefunded, error) {
	event := new(CrowdfundingDonationRefunded)
	if err := _Crowdfunding.contract.UnpackLog(event, "DonationRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdfundingFundsWithdrawnIterator is returned from FilterFundsWithdrawn and is used to iterate over the raw logs and unpacked data for FundsWithdrawn events raised by the Crowdfunding contract.
type CrowdfundingFundsWithdrawnIterator struct {
	Event *CrowdfundingFundsWithdrawn // Event containing the contract specifics and raw log

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
func (it *CrowdfundingFundsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdfundingFundsWithdrawn)
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
		it.Event = new(CrowdfundingFundsWithdrawn)
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
func (it *CrowdfundingFundsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdfundingFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdfundingFundsWithdrawn represents a FundsWithdrawn event raised by the Crowdfunding contract.
type CrowdfundingFundsWithdrawn struct {
	CampaignId *big.Int
	Owner      common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFundsWithdrawn is a free log retrieval operation binding the contract event 0xf440aec6b52895984d061d622e6edeba6210f7c3e059be920663140c084560d7.
//
// Solidity: event FundsWithdrawn(uint256 indexed campaignId, address indexed owner, uint256 amount)
func (_Crowdfunding *CrowdfundingFilterer) FilterFundsWithdrawn(opts *bind.FilterOpts, campaignId []*big.Int, owner []common.Address) (*CrowdfundingFundsWithdrawnIterator, error) {

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Crowdfunding.contract.FilterLogs(opts, "FundsWithdrawn", campaignIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CrowdfundingFundsWithdrawnIterator{contract: _Crowdfunding.contract, event: "FundsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFundsWithdrawn is a free log subscription operation binding the contract event 0xf440aec6b52895984d061d622e6edeba6210f7c3e059be920663140c084560d7.
//
// Solidity: event FundsWithdrawn(uint256 indexed campaignId, address indexed owner, uint256 amount)
func (_Crowdfunding *CrowdfundingFilterer) WatchFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *CrowdfundingFundsWithdrawn, campaignId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Crowdfunding.contract.WatchLogs(opts, "FundsWithdrawn", campaignIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdfundingFundsWithdrawn)
				if err := _Crowdfunding.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
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

// ParseFundsWithdrawn is a log parse operation binding the contract event 0xf440aec6b52895984d061d622e6edeba6210f7c3e059be920663140c084560d7.
//
// Solidity: event FundsWithdrawn(uint256 indexed campaignId, address indexed owner, uint256 amount)
func (_Crowdfunding *CrowdfundingFilterer) ParseFundsWithdrawn(log types.Log) (*CrowdfundingFundsWithdrawn, error) {
	event := new(CrowdfundingFundsWithdrawn)
	if err := _Crowdfunding.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
