package command

import (
	//"crypto"
	"fmt"
	"log"
	"math/big"
	"salary_demo/contract"
	myCrypto "salary_demo/crypto"
	"salary_demo/utils"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var keystoreTest = `
{
  "address": "1a9ec3b0b807464e6d3398a59d6b0a369bf422fa",
  "crypto": {
	"cipher": "aes-128-ctr",
	"ciphertext": "a471054846fb03e3e271339204420806334d1f09d6da40605a1a152e0d8e35f3",
	"cipherparams": {
	  "iv": "44c5095dc698392c55a65aae46e0b5d9"
	},
	"kdf": "scrypt",
	"kdfparams": {
	  "dklen": 32,
	  "n": 262144,
	  "p": 1,
	  "r": 8,
	  "salt": "e0a5fbaecaa3e75e20bccf61ee175141f3597d3b1bae6a28fe09f3507e63545e"
	},
	"mac": "cb3f62975cf6e7dfb454c2973bdd4a59f87262956d5534cdc87fb35703364043"
  },
  "id": "e08301fb-a263-4643-9c2b-d28959f66d6a",
  "version": 3
}
`
var pwdTest = "123"

const generator = 7

const prime = 150001

//random number (1,prime-1)
const privateKey = 113

const network = "http://127.0.0.1:8545"

//Environment includes the backends and contract
type Environment struct {
	Backend  *backends.SimulatedBackend
	Contract *contract.Contract
}

//Keystore represents the keystore and password
type Keystore struct {
	Key string
	Pwd string
}

//Account represents all account information, which includes keystore and private key for elgamal encrption.
type Account struct {
	//ChainKey   *Keystore
	AccountKey *myCrypto.PrivateKey
	Auth       *bind.TransactOpts
}

//InitSim simulate blockchain in local and depoly salary contract automatically.
func InitSim() (env *Environment, account *Account, err error) {
	account = new(Account)

	chainKey := &Keystore{
		Key: keystoreTest,
		Pwd: pwdTest,
	}

	priv := &myCrypto.PrivateKey{
		PublicKey: myCrypto.PublicKey{
			G: new(big.Int).SetInt64(generator),
			P: new(big.Int).SetInt64(prime),
		},
		X: new(big.Int).SetInt64(privateKey),
	}
	priv.H = new(big.Int).Exp(priv.G, priv.X, priv.P)
	account.AccountKey = priv

	//Create a genesis account with a given keystore here
	fmt.Println("create a genesis account...")
	account.Auth, err = bind.NewTransactor(strings.NewReader(chainKey.Key), chainKey.Pwd)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
		return
	}
	addrTest := "0x" + common.Bytes2Hex(account.Auth.From.Bytes())
	fmt.Print("account address: ")
	fmt.Printf(utils.Blue, addrTest)

	//make the genesis account with a bunch of Ether
	alloc := make(core.GenesisAlloc)
	alloc[account.Auth.From] = core.GenesisAccount{Balance: big.NewInt(1337000000000)}

	//initiate the simlated backend instead of using testnet, the default gaslimte is 10000000
	//if gas limite < 10000000, can not deploy contract
	//max gaslimit = 9223372036854775807
	fmt.Println("simulate a blockchain...")
	env = new(Environment)
	env.Backend = backends.NewSimulatedBackend(alloc, 1000000000)

	//deploy the contract which is generated by abigen
	fmt.Println("deploy salary contract...")
	addr, _, sc, err := contract.DeployContract(account.Auth, env.Backend)
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
		return
	}
	env.Contract = sc
	//fmt.Println("The deploy tx hash:", "0x"+common.Bytes2Hex(txDp.Hash().Bytes()))
	fmt.Print("salary contract address: ")
	fmt.Printf(utils.Blue, "0x"+common.Bytes2Hex(addr.Bytes()))
	//fmt.Println("=====Commit all pending txs and fresh state.")
	env.Backend.Commit()

	return
}

//Init loads blockchain such as Ehtereum or Ganache and create a transctor from given private key.
func Init(privKey, contractAddress string) (env *Environment, account *Account, err error) {
	env = new(Environment)
	account = new(Account)
	env.Backend = nil

	//connect to a ethereum network, which is represented as host:port
	conn, err := ethclient.Dial(network)

	fmt.Println("connect to local geth node...", conn)
	if err != nil {
		log.Fatalf("could not connect to local node: %v", err)
	}

	//create a transctor from given private key
	pk, err := crypto.HexToECDSA(privKey)
	auth := bind.NewKeyedTransactor(pk)
	account.Auth = auth
	addrTest := "0x" + common.Bytes2Hex(account.Auth.From.Bytes())
	fmt.Print("account address: ")
	fmt.Printf(utils.Blue, addrTest)

	//deploy the contract if we have not yet deployed it.
	if contractAddress == "" {
		var addr common.Address
		addr, _, env.Contract, err = contract.DeployContract(account.Auth, conn)
		if err != nil {
			log.Fatalf("could not deploy contract: %v", err)
			return
		}
		fmt.Print("salary contract address: ")
		//fmt.Println(utils.Blue, addr)
		fmt.Println("0x" + common.Bytes2Hex(addr.Bytes()))
	} else {
		//load the contract from given address
		env.Contract, err = contract.NewContract(common.HexToAddress(contractAddress), conn)
		if err != nil {
			log.Fatalf("could not get contract: %v", err)
			return
		}
		fmt.Print("salary contract address: ")
		fmt.Println(utils.Blue, contractAddress)
	}

	priv := &myCrypto.PrivateKey{
		PublicKey: myCrypto.PublicKey{
			G: new(big.Int).SetInt64(generator),
			P: new(big.Int).SetInt64(prime),
		},
		X: new(big.Int).SetInt64(privateKey),
	}
	priv.H = new(big.Int).Exp(priv.G, priv.X, priv.P)
	account.AccountKey = priv

	return
}
