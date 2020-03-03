// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_t\",\"type\":\"uint256\"},{\"name\":\"_proof\",\"type\":\"uint256[]\"},{\"name\":\"_cArray\",\"type\":\"uint256[]\"},{\"name\":\"_dArray\",\"type\":\"uint256[]\"}],\"name\":\"takeMoney\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"paySalary\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_h\",\"type\":\"uint256\"},{\"name\":\"_p\",\"type\":\"uint256\"},{\"name\":\"_q\",\"type\":\"uint256\"},{\"name\":\"_mk\",\"type\":\"uint256\"},{\"name\":\"_proof1\",\"type\":\"uint256\"},{\"name\":\"_proof2\",\"type\":\"uint256\"},{\"name\":\"_proof3\",\"type\":\"uint256\"}],\"name\":\"verifyMembershipTest\",\"outputs\":[{\"name\":\"x\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"c\",\"type\":\"uint256\"},{\"name\":\"d\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"salary\",\"outputs\":[{\"name\":\"c\",\"type\":\"uint256\"},{\"name\":\"d\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"checkBalance\",\"outputs\":[{\"name\":\"_c\",\"type\":\"uint256\"},{\"name\":\"_d\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_p\",\"type\":\"uint256\"},{\"name\":\"_q\",\"type\":\"uint256\"},{\"name\":\"_mk\",\"type\":\"uint256\"},{\"name\":\"_proof1\",\"type\":\"uint256\"},{\"name\":\"_proof2\",\"type\":\"uint256\"},{\"name\":\"_proof3\",\"type\":\"uint256\"}],\"name\":\"verifyMembership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_x\",\"type\":\"uint256\"}],\"name\":\"toBytes\",\"outputs\":[{\"name\":\"b\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"generator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_proof\",\"type\":\"uint256[]\"},{\"name\":\"_cArray\",\"type\":\"uint256[]\"},{\"name\":\"_dArray\",\"type\":\"uint256[]\"},{\"name\":\"_c\",\"type\":\"uint256\"},{\"name\":\"_d\",\"type\":\"uint256\"}],\"name\":\"verifyRangeProof\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_c\",\"type\":\"uint256\"},{\"name\":\"_d\",\"type\":\"uint256\"}],\"name\":\"setSalary\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_x\",\"type\":\"uint256\"},{\"name\":\"_y\",\"type\":\"uint256\"}],\"name\":\"myHash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"publicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"checkSalary\",\"outputs\":[{\"name\":\"_c\",\"type\":\"uint256\"},{\"name\":\"_d\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_c\",\"type\":\"uint256\"},{\"name\":\"_d\",\"type\":\"uint256\"},{\"name\":\"_h\",\"type\":\"uint256\"}],\"name\":\"setBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"prime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value1\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_value2\",\"type\":\"uint256\"}],\"name\":\"SetBalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value1\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_value2\",\"type\":\"uint256\"}],\"name\":\"SetSalary\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"PaySalary\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"TakeMoney\",\"type\":\"event\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
var ContractBin = "0x608060405234801561001057600080fd5b5060008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555060008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055506000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001819055506000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055506007600381905550620249f16004819055506118bd806101506000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80637afa1eed11610097578063888ade0a11610066578063888ade0a14610963578063abd73a54146109bb578063bf83496314610a1a578063c7ee005e14610a9457610100565b80637afa1eed1461065d5780637c9095c51461067b5780637ebe3797146108a75780638201f2961461091757610100565b806347b2028d116100d357806347b2028d146104605780635f515226146104bf578063614b493d1461051e578063775a8f5e146105b657610100565b80630b402ebe146101055780630ea37aef1461032757806325e835df1461038357806327e235e314610401575b600080fd5b61030d600480360360a081101561011b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561016257600080fd5b82018360208201111561017457600080fd5b8035906020019184602083028401116401000000008311171561019657600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156101f657600080fd5b82018360208201111561020857600080fd5b8035906020019184602083028401116401000000008311171561022a57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561028a57600080fd5b82018360208201111561029c57600080fd5b803590602001918460208302840111640100000000831117156102be57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290505050610ab2565b604051808215151515815260200191505060405180910390f35b6103696004803603602081101561033d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c9f565b604051808215151515815260200191505060405180910390f35b6103eb600480360360e081101561039957600080fd5b8101908080359060200190929190803590602001909291908035906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190505050610e97565b6040518082815260200191505060405180910390f35b6104436004803603602081101561041757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610fe5565b604051808381526020018281526020019250505060405180910390f35b6104a26004803603602081101561047657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611009565b604051808381526020018281526020019250505060405180910390f35b610501600480360360208110156104d557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061102d565b604051808381526020018281526020019250505060405180910390f35b61059c600480360360e081101561053457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190803590602001909291905050506110bd565b604051808215151515815260200191505060405180910390f35b6105e2600480360360208110156105cc57600080fd5b8101908080359060200190929190505050611299565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610622578082015181840152602081019050610607565b50505050905090810190601f16801561064f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610665611345565b6040518082815260200191505060405180910390f35b61088d600480360360c081101561069157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156106ce57600080fd5b8201836020820111156106e057600080fd5b8035906020019184602083028401116401000000008311171561070257600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561076257600080fd5b82018360208201111561077457600080fd5b8035906020019184602083028401116401000000008311171561079657600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156107f657600080fd5b82018360208201111561080857600080fd5b8035906020019184602083028401116401000000008311171561082a57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290803590602001909291908035906020019092919050505061134b565b604051808215151515815260200191505060405180910390f35b6108fd600480360360608110156108bd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919050505061148c565b604051808215151515815260200191505060405180910390f35b61094d6004803603604081101561092d57600080fd5b81019080803590602001909291908035906020019092919050505061157d565b6040518082815260200191505060405180910390f35b6109a56004803603602081101561097957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061158a565b6040518082815260200191505060405180910390f35b6109fd600480360360208110156109d157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506115a2565b604051808381526020018281526020019250505060405180910390f35b610a7a60048036036080811015610a3057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919080359060200190929190505050611634565b604051808215151515815260200191505060405180910390f35b610a9c611768565b6040518082815260200191505060405180910390f35b6000808511610b29576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f4d6f6e65792073686f756c64206265206120706f7369746976652e000000000081525060200191505060405180910390fd5b60008060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015490506000610b8b610b836003548960045461176e565b600454611870565b905060045480610b9757fe5b816000808b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154099050610beb88878787868661134b565b15610c8f57806000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055508773ffffffffffffffffffffffffffffffffffffffff167fda55bc161f61953c91447a146006ad4921cc089b3b3ebc716664ef9e47266f27886040518082815260200191505060405180910390a2600192505050610c96565b6000925050505b95945050505050565b600060045480610cab57fe5b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001546000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154096000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555060045480610d8057fe5b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101546000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154096000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055508173ffffffffffffffffffffffffffffffffffffffff167fd7d07c6746242e3a94807f5fcdfe276b7bc086bfaf8ebc45a9ba69dcadc0ea5960405160405180910390a260019050919050565b600080610eb2610eaa8a8760045461176e565b600454611870565b90506000610ec56003548660045461176e565b905060045480610ed157fe5b82820990506000610ee58a8760045461176e565b9050610ef4898860045461176e565b9250610f0283600454611870565b925060045480610f0e57fe5b83820990506000610f1f838361157d565b9050610f39610f318d8360045461176e565b600454611870565b93506000610f4c6003548860045461176e565b905060045480610f5857fe5b85820990506000610f6c8d8960045461176e565b9050610f7d6003548c60045461176e565b9550610f8b86600454611870565b955060045480610f9757fe5b868d099550610fa9868460045461176e565b9550610fb786600454611870565b955060045480610fc357fe5b8682099050610fd2828261157d565b9650505050505050979650505050505050565b60006020528060005260406000206000915090508060000154908060010154905082565b60016020528060005260406000206000915090508060000154908060010154905082565b6000806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001546000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015491509150915091565b60008061111761110f600260008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548760045461176e565b600454611870565b9050600061112a6003548660045461176e565b90506004548061113657fe5b8282099050600061114a8a8760045461176e565b9050611159898860045461176e565b925061116783600454611870565b92506004548061117357fe5b83820990506000611184838361157d565b90506111dd6111d5600260008f73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548360045461176e565b600454611870565b935060006111f06003548860045461176e565b9050600454806111fc57fe5b858209905060006112108d8960045461176e565b90506112216003548c60045461176e565b955061122f86600454611870565b95506004548061123b57fe5b868d09955061124d868460045461176e565b955061125b86600454611870565b95506004548061126757fe5b8682099050611276828261157d565b8a14611283576000611286565b60015b9650505050505050979650505050505050565b606060206040519080825280601f01601f1916602001820160405280156112cf5781602001600182028038833980820191505090505b50905060008090505b602081101561133f5780601f0360080260020a83816112f357fe5b0460f81b82828151811061130357fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806001019150506112d8565b50919050565b60035481565b600080600190506000600190506001925060008090505b601f8111611444576004548061137457fe5b88828151811061138057fe5b6020026020010151840992506004548061139657fe5b8782815181106113a257fe5b60200260200101518309915083801561143757506114368a8983815181106113c657fe5b60200260200101518984815181106113da57fe5b60200260200101518460020a8d60038702815181106113f557fe5b60200260200101518e600160038902018151811061140f57fe5b60200260200101518f600260038a02018151811061142957fe5b60200260200101516110bd565b5b9350600181019050611362565b5082801561145f575081851461145b57600061145e565b60015b5b925082801561147b575080841461147757600061147a565b60015b5b925082925050509695505050505050565b600082600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555081600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055508373ffffffffffffffffffffffffffffffffffffffff167f70c22001aa542796dc91ca19cb337e7ea7d4fcf47d2b7ddcbc0ee3fd856c74798484604051808381526020018281526020019250505060405180910390a2600190509392505050565b6000818301905092915050565b60026020528060005260406000206000915090505481565b600080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015491509150915091565b6000836000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000181905550826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018190555081600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508473ffffffffffffffffffffffffffffffffffffffff167f301faec51960184067a8ff5dc1efc65fa6e11942352ff244bb4b3cecb5de8e3c8585604051808381526020018281526020019250505060405180910390a260019050949350505050565b60045481565b60008084118015611780575060008310155b801561178c5750600082115b6117fe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f782c792c7a2073686f756c6420626520706f736974697665732e00000000000081525060200191505060405180910390fd5b60008314156118105760019050611869565b600190506000849050600084905060008490506000600190505b8281116118615760008184161461184957818061184357fe5b84860994505b818061185157fe5b848509935060028102905061182a565b508393505050505b9392505050565b600061188083600284038461176e565b90509291505056fea265627a7a72305820d1c62177a6e58b97ba54888089556871f4b4ff982f874ac6516471694210adec64736f6c634300050a0032"

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256 c, uint256 d)
func (_Contract *ContractCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (struct {
	C *big.Int
	D *big.Int
}, error) {
	ret := new(struct {
		C *big.Int
		D *big.Int
	})
	out := ret
	err := _Contract.contract.Call(opts, out, "balances", arg0)
	return *ret, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256 c, uint256 d)
func (_Contract *ContractSession) Balances(arg0 common.Address) (struct {
	C *big.Int
	D *big.Int
}, error) {
	return _Contract.Contract.Balances(&_Contract.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256 c, uint256 d)
func (_Contract *ContractCallerSession) Balances(arg0 common.Address) (struct {
	C *big.Int
	D *big.Int
}, error) {
	return _Contract.Contract.Balances(&_Contract.CallOpts, arg0)
}

// CheckBalance is a free data retrieval call binding the contract method 0x5f515226.
//
// Solidity: function checkBalance(address _addr) constant returns(uint256 _c, uint256 _d)
func (_Contract *ContractCaller) CheckBalance(opts *bind.CallOpts, _addr common.Address) (struct {
	C *big.Int
	D *big.Int
}, error) {
	ret := new(struct {
		C *big.Int
		D *big.Int
	})
	out := ret
	err := _Contract.contract.Call(opts, out, "checkBalance", _addr)
	return *ret, err
}

// CheckBalance is a free data retrieval call binding the contract method 0x5f515226.
//
// Solidity: function checkBalance(address _addr) constant returns(uint256 _c, uint256 _d)
func (_Contract *ContractSession) CheckBalance(_addr common.Address) (struct {
	C *big.Int
	D *big.Int
}, error) {
	return _Contract.Contract.CheckBalance(&_Contract.CallOpts, _addr)
}

// CheckBalance is a free data retrieval call binding the contract method 0x5f515226.
//
// Solidity: function checkBalance(address _addr) constant returns(uint256 _c, uint256 _d)
func (_Contract *ContractCallerSession) CheckBalance(_addr common.Address) (struct {
	C *big.Int
	D *big.Int
}, error) {
	return _Contract.Contract.CheckBalance(&_Contract.CallOpts, _addr)
}

// CheckSalary is a free data retrieval call binding the contract method 0xabd73a54.
//
// Solidity: function checkSalary(address _addr) constant returns(uint256 _c, uint256 _d)
func (_Contract *ContractCaller) CheckSalary(opts *bind.CallOpts, _addr common.Address) (struct {
	C *big.Int
	D *big.Int
}, error) {
	ret := new(struct {
		C *big.Int
		D *big.Int
	})
	out := ret
	err := _Contract.contract.Call(opts, out, "checkSalary", _addr)
	return *ret, err
}

// CheckSalary is a free data retrieval call binding the contract method 0xabd73a54.
//
// Solidity: function checkSalary(address _addr) constant returns(uint256 _c, uint256 _d)
func (_Contract *ContractSession) CheckSalary(_addr common.Address) (struct {
	C *big.Int
	D *big.Int
}, error) {
	return _Contract.Contract.CheckSalary(&_Contract.CallOpts, _addr)
}

// CheckSalary is a free data retrieval call binding the contract method 0xabd73a54.
//
// Solidity: function checkSalary(address _addr) constant returns(uint256 _c, uint256 _d)
func (_Contract *ContractCallerSession) CheckSalary(_addr common.Address) (struct {
	C *big.Int
	D *big.Int
}, error) {
	return _Contract.Contract.CheckSalary(&_Contract.CallOpts, _addr)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() constant returns(uint256)
func (_Contract *ContractCaller) Generator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "generator")
	return *ret0, err
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() constant returns(uint256)
func (_Contract *ContractSession) Generator() (*big.Int, error) {
	return _Contract.Contract.Generator(&_Contract.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() constant returns(uint256)
func (_Contract *ContractCallerSession) Generator() (*big.Int, error) {
	return _Contract.Contract.Generator(&_Contract.CallOpts)
}

// MyHash is a free data retrieval call binding the contract method 0x8201f296.
//
// Solidity: function myHash(uint256 _x, uint256 _y) constant returns(uint256)
func (_Contract *ContractCaller) MyHash(opts *bind.CallOpts, _x *big.Int, _y *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "myHash", _x, _y)
	return *ret0, err
}

// MyHash is a free data retrieval call binding the contract method 0x8201f296.
//
// Solidity: function myHash(uint256 _x, uint256 _y) constant returns(uint256)
func (_Contract *ContractSession) MyHash(_x *big.Int, _y *big.Int) (*big.Int, error) {
	return _Contract.Contract.MyHash(&_Contract.CallOpts, _x, _y)
}

// MyHash is a free data retrieval call binding the contract method 0x8201f296.
//
// Solidity: function myHash(uint256 _x, uint256 _y) constant returns(uint256)
func (_Contract *ContractCallerSession) MyHash(_x *big.Int, _y *big.Int) (*big.Int, error) {
	return _Contract.Contract.MyHash(&_Contract.CallOpts, _x, _y)
}

// Prime is a free data retrieval call binding the contract method 0xc7ee005e.
//
// Solidity: function prime() constant returns(uint256)
func (_Contract *ContractCaller) Prime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "prime")
	return *ret0, err
}

// Prime is a free data retrieval call binding the contract method 0xc7ee005e.
//
// Solidity: function prime() constant returns(uint256)
func (_Contract *ContractSession) Prime() (*big.Int, error) {
	return _Contract.Contract.Prime(&_Contract.CallOpts)
}

// Prime is a free data retrieval call binding the contract method 0xc7ee005e.
//
// Solidity: function prime() constant returns(uint256)
func (_Contract *ContractCallerSession) Prime() (*big.Int, error) {
	return _Contract.Contract.Prime(&_Contract.CallOpts)
}

// PublicKey is a free data retrieval call binding the contract method 0x888ade0a.
//
// Solidity: function publicKey(address ) constant returns(uint256)
func (_Contract *ContractCaller) PublicKey(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "publicKey", arg0)
	return *ret0, err
}

// PublicKey is a free data retrieval call binding the contract method 0x888ade0a.
//
// Solidity: function publicKey(address ) constant returns(uint256)
func (_Contract *ContractSession) PublicKey(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.PublicKey(&_Contract.CallOpts, arg0)
}

// PublicKey is a free data retrieval call binding the contract method 0x888ade0a.
//
// Solidity: function publicKey(address ) constant returns(uint256)
func (_Contract *ContractCallerSession) PublicKey(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.PublicKey(&_Contract.CallOpts, arg0)
}

// Salary is a free data retrieval call binding the contract method 0x47b2028d.
//
// Solidity: function salary(address ) constant returns(uint256 c, uint256 d)
func (_Contract *ContractCaller) Salary(opts *bind.CallOpts, arg0 common.Address) (struct {
	C *big.Int
	D *big.Int
}, error) {
	ret := new(struct {
		C *big.Int
		D *big.Int
	})
	out := ret
	err := _Contract.contract.Call(opts, out, "salary", arg0)
	return *ret, err
}

// Salary is a free data retrieval call binding the contract method 0x47b2028d.
//
// Solidity: function salary(address ) constant returns(uint256 c, uint256 d)
func (_Contract *ContractSession) Salary(arg0 common.Address) (struct {
	C *big.Int
	D *big.Int
}, error) {
	return _Contract.Contract.Salary(&_Contract.CallOpts, arg0)
}

// Salary is a free data retrieval call binding the contract method 0x47b2028d.
//
// Solidity: function salary(address ) constant returns(uint256 c, uint256 d)
func (_Contract *ContractCallerSession) Salary(arg0 common.Address) (struct {
	C *big.Int
	D *big.Int
}, error) {
	return _Contract.Contract.Salary(&_Contract.CallOpts, arg0)
}

// ToBytes is a free data retrieval call binding the contract method 0x775a8f5e.
//
// Solidity: function toBytes(uint256 _x) constant returns(bytes b)
func (_Contract *ContractCaller) ToBytes(opts *bind.CallOpts, _x *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "toBytes", _x)
	return *ret0, err
}

// ToBytes is a free data retrieval call binding the contract method 0x775a8f5e.
//
// Solidity: function toBytes(uint256 _x) constant returns(bytes b)
func (_Contract *ContractSession) ToBytes(_x *big.Int) ([]byte, error) {
	return _Contract.Contract.ToBytes(&_Contract.CallOpts, _x)
}

// ToBytes is a free data retrieval call binding the contract method 0x775a8f5e.
//
// Solidity: function toBytes(uint256 _x) constant returns(bytes b)
func (_Contract *ContractCallerSession) ToBytes(_x *big.Int) ([]byte, error) {
	return _Contract.Contract.ToBytes(&_Contract.CallOpts, _x)
}

// VerifyMembership is a free data retrieval call binding the contract method 0x614b493d.
//
// Solidity: function verifyMembership(address _addr, uint256 _p, uint256 _q, uint256 _mk, uint256 _proof1, uint256 _proof2, uint256 _proof3) constant returns(bool)
func (_Contract *ContractCaller) VerifyMembership(opts *bind.CallOpts, _addr common.Address, _p *big.Int, _q *big.Int, _mk *big.Int, _proof1 *big.Int, _proof2 *big.Int, _proof3 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "verifyMembership", _addr, _p, _q, _mk, _proof1, _proof2, _proof3)
	return *ret0, err
}

// VerifyMembership is a free data retrieval call binding the contract method 0x614b493d.
//
// Solidity: function verifyMembership(address _addr, uint256 _p, uint256 _q, uint256 _mk, uint256 _proof1, uint256 _proof2, uint256 _proof3) constant returns(bool)
func (_Contract *ContractSession) VerifyMembership(_addr common.Address, _p *big.Int, _q *big.Int, _mk *big.Int, _proof1 *big.Int, _proof2 *big.Int, _proof3 *big.Int) (bool, error) {
	return _Contract.Contract.VerifyMembership(&_Contract.CallOpts, _addr, _p, _q, _mk, _proof1, _proof2, _proof3)
}

// VerifyMembership is a free data retrieval call binding the contract method 0x614b493d.
//
// Solidity: function verifyMembership(address _addr, uint256 _p, uint256 _q, uint256 _mk, uint256 _proof1, uint256 _proof2, uint256 _proof3) constant returns(bool)
func (_Contract *ContractCallerSession) VerifyMembership(_addr common.Address, _p *big.Int, _q *big.Int, _mk *big.Int, _proof1 *big.Int, _proof2 *big.Int, _proof3 *big.Int) (bool, error) {
	return _Contract.Contract.VerifyMembership(&_Contract.CallOpts, _addr, _p, _q, _mk, _proof1, _proof2, _proof3)
}

// VerifyMembershipTest is a free data retrieval call binding the contract method 0x25e835df.
//
// Solidity: function verifyMembershipTest(uint256 _h, uint256 _p, uint256 _q, uint256 _mk, uint256 _proof1, uint256 _proof2, uint256 _proof3) constant returns(uint256 x)
func (_Contract *ContractCaller) VerifyMembershipTest(opts *bind.CallOpts, _h *big.Int, _p *big.Int, _q *big.Int, _mk *big.Int, _proof1 *big.Int, _proof2 *big.Int, _proof3 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "verifyMembershipTest", _h, _p, _q, _mk, _proof1, _proof2, _proof3)
	return *ret0, err
}

// VerifyMembershipTest is a free data retrieval call binding the contract method 0x25e835df.
//
// Solidity: function verifyMembershipTest(uint256 _h, uint256 _p, uint256 _q, uint256 _mk, uint256 _proof1, uint256 _proof2, uint256 _proof3) constant returns(uint256 x)
func (_Contract *ContractSession) VerifyMembershipTest(_h *big.Int, _p *big.Int, _q *big.Int, _mk *big.Int, _proof1 *big.Int, _proof2 *big.Int, _proof3 *big.Int) (*big.Int, error) {
	return _Contract.Contract.VerifyMembershipTest(&_Contract.CallOpts, _h, _p, _q, _mk, _proof1, _proof2, _proof3)
}

// VerifyMembershipTest is a free data retrieval call binding the contract method 0x25e835df.
//
// Solidity: function verifyMembershipTest(uint256 _h, uint256 _p, uint256 _q, uint256 _mk, uint256 _proof1, uint256 _proof2, uint256 _proof3) constant returns(uint256 x)
func (_Contract *ContractCallerSession) VerifyMembershipTest(_h *big.Int, _p *big.Int, _q *big.Int, _mk *big.Int, _proof1 *big.Int, _proof2 *big.Int, _proof3 *big.Int) (*big.Int, error) {
	return _Contract.Contract.VerifyMembershipTest(&_Contract.CallOpts, _h, _p, _q, _mk, _proof1, _proof2, _proof3)
}

// VerifyRangeProof is a free data retrieval call binding the contract method 0x7c9095c5.
//
// Solidity: function verifyRangeProof(address _addr, uint256[] _proof, uint256[] _cArray, uint256[] _dArray, uint256 _c, uint256 _d) constant returns(bool result)
func (_Contract *ContractCaller) VerifyRangeProof(opts *bind.CallOpts, _addr common.Address, _proof []*big.Int, _cArray []*big.Int, _dArray []*big.Int, _c *big.Int, _d *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "verifyRangeProof", _addr, _proof, _cArray, _dArray, _c, _d)
	return *ret0, err
}

// VerifyRangeProof is a free data retrieval call binding the contract method 0x7c9095c5.
//
// Solidity: function verifyRangeProof(address _addr, uint256[] _proof, uint256[] _cArray, uint256[] _dArray, uint256 _c, uint256 _d) constant returns(bool result)
func (_Contract *ContractSession) VerifyRangeProof(_addr common.Address, _proof []*big.Int, _cArray []*big.Int, _dArray []*big.Int, _c *big.Int, _d *big.Int) (bool, error) {
	return _Contract.Contract.VerifyRangeProof(&_Contract.CallOpts, _addr, _proof, _cArray, _dArray, _c, _d)
}

// VerifyRangeProof is a free data retrieval call binding the contract method 0x7c9095c5.
//
// Solidity: function verifyRangeProof(address _addr, uint256[] _proof, uint256[] _cArray, uint256[] _dArray, uint256 _c, uint256 _d) constant returns(bool result)
func (_Contract *ContractCallerSession) VerifyRangeProof(_addr common.Address, _proof []*big.Int, _cArray []*big.Int, _dArray []*big.Int, _c *big.Int, _d *big.Int) (bool, error) {
	return _Contract.Contract.VerifyRangeProof(&_Contract.CallOpts, _addr, _proof, _cArray, _dArray, _c, _d)
}

// PaySalary is a paid mutator transaction binding the contract method 0x0ea37aef.
//
// Solidity: function paySalary(address _addr) returns(bool)
func (_Contract *ContractTransactor) PaySalary(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "paySalary", _addr)
}

// PaySalary is a paid mutator transaction binding the contract method 0x0ea37aef.
//
// Solidity: function paySalary(address _addr) returns(bool)
func (_Contract *ContractSession) PaySalary(_addr common.Address) (*types.Transaction, error) {
	return _Contract.Contract.PaySalary(&_Contract.TransactOpts, _addr)
}

// PaySalary is a paid mutator transaction binding the contract method 0x0ea37aef.
//
// Solidity: function paySalary(address _addr) returns(bool)
func (_Contract *ContractTransactorSession) PaySalary(_addr common.Address) (*types.Transaction, error) {
	return _Contract.Contract.PaySalary(&_Contract.TransactOpts, _addr)
}

// SetBalance is a paid mutator transaction binding the contract method 0xbf834963.
//
// Solidity: function setBalance(address _addr, uint256 _c, uint256 _d, uint256 _h) returns(bool)
func (_Contract *ContractTransactor) SetBalance(opts *bind.TransactOpts, _addr common.Address, _c *big.Int, _d *big.Int, _h *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setBalance", _addr, _c, _d, _h)
}

// SetBalance is a paid mutator transaction binding the contract method 0xbf834963.
//
// Solidity: function setBalance(address _addr, uint256 _c, uint256 _d, uint256 _h) returns(bool)
func (_Contract *ContractSession) SetBalance(_addr common.Address, _c *big.Int, _d *big.Int, _h *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetBalance(&_Contract.TransactOpts, _addr, _c, _d, _h)
}

// SetBalance is a paid mutator transaction binding the contract method 0xbf834963.
//
// Solidity: function setBalance(address _addr, uint256 _c, uint256 _d, uint256 _h) returns(bool)
func (_Contract *ContractTransactorSession) SetBalance(_addr common.Address, _c *big.Int, _d *big.Int, _h *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetBalance(&_Contract.TransactOpts, _addr, _c, _d, _h)
}

// SetSalary is a paid mutator transaction binding the contract method 0x7ebe3797.
//
// Solidity: function setSalary(address _addr, uint256 _c, uint256 _d) returns(bool)
func (_Contract *ContractTransactor) SetSalary(opts *bind.TransactOpts, _addr common.Address, _c *big.Int, _d *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setSalary", _addr, _c, _d)
}

// SetSalary is a paid mutator transaction binding the contract method 0x7ebe3797.
//
// Solidity: function setSalary(address _addr, uint256 _c, uint256 _d) returns(bool)
func (_Contract *ContractSession) SetSalary(_addr common.Address, _c *big.Int, _d *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetSalary(&_Contract.TransactOpts, _addr, _c, _d)
}

// SetSalary is a paid mutator transaction binding the contract method 0x7ebe3797.
//
// Solidity: function setSalary(address _addr, uint256 _c, uint256 _d) returns(bool)
func (_Contract *ContractTransactorSession) SetSalary(_addr common.Address, _c *big.Int, _d *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetSalary(&_Contract.TransactOpts, _addr, _c, _d)
}

// TakeMoney is a paid mutator transaction binding the contract method 0x0b402ebe.
//
// Solidity: function takeMoney(address _addr, uint256 _t, uint256[] _proof, uint256[] _cArray, uint256[] _dArray) returns(bool)
func (_Contract *ContractTransactor) TakeMoney(opts *bind.TransactOpts, _addr common.Address, _t *big.Int, _proof []*big.Int, _cArray []*big.Int, _dArray []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "takeMoney", _addr, _t, _proof, _cArray, _dArray)
}

// TakeMoney is a paid mutator transaction binding the contract method 0x0b402ebe.
//
// Solidity: function takeMoney(address _addr, uint256 _t, uint256[] _proof, uint256[] _cArray, uint256[] _dArray) returns(bool)
func (_Contract *ContractSession) TakeMoney(_addr common.Address, _t *big.Int, _proof []*big.Int, _cArray []*big.Int, _dArray []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TakeMoney(&_Contract.TransactOpts, _addr, _t, _proof, _cArray, _dArray)
}

// TakeMoney is a paid mutator transaction binding the contract method 0x0b402ebe.
//
// Solidity: function takeMoney(address _addr, uint256 _t, uint256[] _proof, uint256[] _cArray, uint256[] _dArray) returns(bool)
func (_Contract *ContractTransactorSession) TakeMoney(_addr common.Address, _t *big.Int, _proof []*big.Int, _cArray []*big.Int, _dArray []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TakeMoney(&_Contract.TransactOpts, _addr, _t, _proof, _cArray, _dArray)
}

// ContractPaySalaryIterator is returned from FilterPaySalary and is used to iterate over the raw logs and unpacked data for PaySalary events raised by the Contract contract.
type ContractPaySalaryIterator struct {
	Event *ContractPaySalary // Event containing the contract specifics and raw log

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
func (it *ContractPaySalaryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPaySalary)
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
		it.Event = new(ContractPaySalary)
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
func (it *ContractPaySalaryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPaySalaryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPaySalary represents a PaySalary event raised by the Contract contract.
type ContractPaySalary struct {
	To  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaySalary is a free log retrieval operation binding the contract event 0xd7d07c6746242e3a94807f5fcdfe276b7bc086bfaf8ebc45a9ba69dcadc0ea59.
//
// Solidity: event PaySalary(address indexed _to)
func (_Contract *ContractFilterer) FilterPaySalary(opts *bind.FilterOpts, _to []common.Address) (*ContractPaySalaryIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PaySalary", _toRule)
	if err != nil {
		return nil, err
	}
	return &ContractPaySalaryIterator{contract: _Contract.contract, event: "PaySalary", logs: logs, sub: sub}, nil
}

// WatchPaySalary is a free log subscription operation binding the contract event 0xd7d07c6746242e3a94807f5fcdfe276b7bc086bfaf8ebc45a9ba69dcadc0ea59.
//
// Solidity: event PaySalary(address indexed _to)
func (_Contract *ContractFilterer) WatchPaySalary(opts *bind.WatchOpts, sink chan<- *ContractPaySalary, _to []common.Address) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PaySalary", _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPaySalary)
				if err := _Contract.contract.UnpackLog(event, "PaySalary", log); err != nil {
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

// ParsePaySalary is a log parse operation binding the contract event 0xd7d07c6746242e3a94807f5fcdfe276b7bc086bfaf8ebc45a9ba69dcadc0ea59.
//
// Solidity: event PaySalary(address indexed _to)
func (_Contract *ContractFilterer) ParsePaySalary(log types.Log) (*ContractPaySalary, error) {
	event := new(ContractPaySalary)
	if err := _Contract.contract.UnpackLog(event, "PaySalary", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractSetBalanceIterator is returned from FilterSetBalance and is used to iterate over the raw logs and unpacked data for SetBalance events raised by the Contract contract.
type ContractSetBalanceIterator struct {
	Event *ContractSetBalance // Event containing the contract specifics and raw log

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
func (it *ContractSetBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetBalance)
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
		it.Event = new(ContractSetBalance)
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
func (it *ContractSetBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetBalance represents a SetBalance event raised by the Contract contract.
type ContractSetBalance struct {
	To     common.Address
	Value1 *big.Int
	Value2 *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetBalance is a free log retrieval operation binding the contract event 0x301faec51960184067a8ff5dc1efc65fa6e11942352ff244bb4b3cecb5de8e3c.
//
// Solidity: event SetBalance(address indexed _to, uint256 _value1, uint256 _value2)
func (_Contract *ContractFilterer) FilterSetBalance(opts *bind.FilterOpts, _to []common.Address) (*ContractSetBalanceIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetBalance", _toRule)
	if err != nil {
		return nil, err
	}
	return &ContractSetBalanceIterator{contract: _Contract.contract, event: "SetBalance", logs: logs, sub: sub}, nil
}

// WatchSetBalance is a free log subscription operation binding the contract event 0x301faec51960184067a8ff5dc1efc65fa6e11942352ff244bb4b3cecb5de8e3c.
//
// Solidity: event SetBalance(address indexed _to, uint256 _value1, uint256 _value2)
func (_Contract *ContractFilterer) WatchSetBalance(opts *bind.WatchOpts, sink chan<- *ContractSetBalance, _to []common.Address) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetBalance", _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetBalance)
				if err := _Contract.contract.UnpackLog(event, "SetBalance", log); err != nil {
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

// ParseSetBalance is a log parse operation binding the contract event 0x301faec51960184067a8ff5dc1efc65fa6e11942352ff244bb4b3cecb5de8e3c.
//
// Solidity: event SetBalance(address indexed _to, uint256 _value1, uint256 _value2)
func (_Contract *ContractFilterer) ParseSetBalance(log types.Log) (*ContractSetBalance, error) {
	event := new(ContractSetBalance)
	if err := _Contract.contract.UnpackLog(event, "SetBalance", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractSetSalaryIterator is returned from FilterSetSalary and is used to iterate over the raw logs and unpacked data for SetSalary events raised by the Contract contract.
type ContractSetSalaryIterator struct {
	Event *ContractSetSalary // Event containing the contract specifics and raw log

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
func (it *ContractSetSalaryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetSalary)
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
		it.Event = new(ContractSetSalary)
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
func (it *ContractSetSalaryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetSalaryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetSalary represents a SetSalary event raised by the Contract contract.
type ContractSetSalary struct {
	To     common.Address
	Value1 *big.Int
	Value2 *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetSalary is a free log retrieval operation binding the contract event 0x70c22001aa542796dc91ca19cb337e7ea7d4fcf47d2b7ddcbc0ee3fd856c7479.
//
// Solidity: event SetSalary(address indexed _to, uint256 _value1, uint256 _value2)
func (_Contract *ContractFilterer) FilterSetSalary(opts *bind.FilterOpts, _to []common.Address) (*ContractSetSalaryIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetSalary", _toRule)
	if err != nil {
		return nil, err
	}
	return &ContractSetSalaryIterator{contract: _Contract.contract, event: "SetSalary", logs: logs, sub: sub}, nil
}

// WatchSetSalary is a free log subscription operation binding the contract event 0x70c22001aa542796dc91ca19cb337e7ea7d4fcf47d2b7ddcbc0ee3fd856c7479.
//
// Solidity: event SetSalary(address indexed _to, uint256 _value1, uint256 _value2)
func (_Contract *ContractFilterer) WatchSetSalary(opts *bind.WatchOpts, sink chan<- *ContractSetSalary, _to []common.Address) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetSalary", _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetSalary)
				if err := _Contract.contract.UnpackLog(event, "SetSalary", log); err != nil {
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

// ParseSetSalary is a log parse operation binding the contract event 0x70c22001aa542796dc91ca19cb337e7ea7d4fcf47d2b7ddcbc0ee3fd856c7479.
//
// Solidity: event SetSalary(address indexed _to, uint256 _value1, uint256 _value2)
func (_Contract *ContractFilterer) ParseSetSalary(log types.Log) (*ContractSetSalary, error) {
	event := new(ContractSetSalary)
	if err := _Contract.contract.UnpackLog(event, "SetSalary", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractTakeMoneyIterator is returned from FilterTakeMoney and is used to iterate over the raw logs and unpacked data for TakeMoney events raised by the Contract contract.
type ContractTakeMoneyIterator struct {
	Event *ContractTakeMoney // Event containing the contract specifics and raw log

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
func (it *ContractTakeMoneyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTakeMoney)
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
		it.Event = new(ContractTakeMoney)
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
func (it *ContractTakeMoneyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTakeMoneyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTakeMoney represents a TakeMoney event raised by the Contract contract.
type ContractTakeMoney struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTakeMoney is a free log retrieval operation binding the contract event 0xda55bc161f61953c91447a146006ad4921cc089b3b3ebc716664ef9e47266f27.
//
// Solidity: event TakeMoney(address indexed _to, uint256 _value)
func (_Contract *ContractFilterer) FilterTakeMoney(opts *bind.FilterOpts, _to []common.Address) (*ContractTakeMoneyIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TakeMoney", _toRule)
	if err != nil {
		return nil, err
	}
	return &ContractTakeMoneyIterator{contract: _Contract.contract, event: "TakeMoney", logs: logs, sub: sub}, nil
}

// WatchTakeMoney is a free log subscription operation binding the contract event 0xda55bc161f61953c91447a146006ad4921cc089b3b3ebc716664ef9e47266f27.
//
// Solidity: event TakeMoney(address indexed _to, uint256 _value)
func (_Contract *ContractFilterer) WatchTakeMoney(opts *bind.WatchOpts, sink chan<- *ContractTakeMoney, _to []common.Address) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TakeMoney", _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTakeMoney)
				if err := _Contract.contract.UnpackLog(event, "TakeMoney", log); err != nil {
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

// ParseTakeMoney is a log parse operation binding the contract event 0xda55bc161f61953c91447a146006ad4921cc089b3b3ebc716664ef9e47266f27.
//
// Solidity: event TakeMoney(address indexed _to, uint256 _value)
func (_Contract *ContractFilterer) ParseTakeMoney(log types.Log) (*ContractTakeMoney, error) {
	event := new(ContractTakeMoney)
	if err := _Contract.contract.UnpackLog(event, "TakeMoney", log); err != nil {
		return nil, err
	}
	return event, nil
}
