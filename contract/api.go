package contract

import (
	"bytes"
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	myCrypto "salary_demo/crypto"

	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

//const ChainNetwork = "http://127.0.0.1:7545"
//const contractAddr = "0xB4e4110Bc0166B56bDa89283e88F1564FcFCeBEB"

//SetBalance sets the balance of a address to a certain value.
func SetBalance(sim *backends.SimulatedBackend, auth *bind.TransactOpts, contract *Contract, account string, amount *big.Int, priv *myCrypto.PrivateKey) (bool, error) {
	C, D, err := myCrypto.EncryptInt(rand.Reader, &priv.PublicKey, amount)
	if err != nil {
		log.Fatalf("could not encrypt balance: %v", err)
		return false, err
	}

	_, err = contract.SetBalance(auth, common.HexToAddress(account), C, D, priv.H)
	if err != nil {
		log.Fatalf("could not set balance: %v", err)
		return false, err
	}
	fmt.Println("encrypted balance:", C, D)
	if sim != nil {
		sim.Commit()
	}
	return true, nil
}

//GetBalance checkes the balance of a address.
func GetBalance(contract *Contract, account string, priv *myCrypto.PrivateKey) (balance *big.Int, err error) {

	balanceEncrypted, err := contract.CheckBalance(nil, common.HexToAddress(account))
	if err != nil {
		log.Fatalf("could not get balance: %v", err)
		return nil, err
	}
	//fmt.Println("log: the c,d of balance:", balanceEncrypted.C, balanceEncrypted.D)
	zero := big.NewInt(0)
	if (balanceEncrypted.C.Cmp(zero) == 0) && (balanceEncrypted.D.Cmp(zero) == 0) {
		return zero, nil
	}
	//No commit here, since there is no need to change chain state

	balance, err = myCrypto.DecryptInt(priv, balanceEncrypted.C, balanceEncrypted.D)
	if err != nil {
		log.Fatalf("could not decrypt balance: %v", err)
		return nil, err
	}
	//fmt.Println("=====Function GetBalance:", account, balance)
	return balance, nil
}

//SetSalary sets the salary of a address to a certain value, which similar to SetBalance
func SetSalary(sim *backends.SimulatedBackend, auth *bind.TransactOpts, contract *Contract, account string, amount *big.Int, priv *myCrypto.PrivateKey) (bool, error) {
	//fmt.Println("=====Function SetSalary:", account, amount)

	C, D, err := myCrypto.EncryptInt(rand.Reader, &priv.PublicKey, amount)
	if err != nil {
		log.Fatalf("could not encrypt salary: %v", err)
		return false, err
	}

	_, err = contract.SetSalary(auth, common.HexToAddress(account), C, D)
	if err != nil {
		log.Fatalf("could not set salary: %v", err)
		return false, err
	}
	fmt.Println("encrypted salary:", C, D)
	sim.Commit()
	return true, nil
}

//GetSalary checkes the salary of a address.
func GetSalary(contract *Contract, account string, priv *myCrypto.PrivateKey) (salary *big.Int, err error) {
	//fmt.Println("=====check the salary of", account)
	salaryEncrypted, err := contract.CheckSalary(nil, common.HexToAddress(account))
	if err != nil {
		log.Fatalf("could not get salary: %v", err)
		return nil, err
	}
	//fmt.Println("log: the c,d of salary:", salaryEncrypted.C, salaryEncrypted.D)
	zero := big.NewInt(0)
	if (salaryEncrypted.C.Cmp(zero) == 0) && (salaryEncrypted.D.Cmp(zero) == 0) {
		return zero, nil
	}
	//No commit here, since there is no need to change chain state

	salary, err = myCrypto.DecryptInt(priv, salaryEncrypted.C, salaryEncrypted.D)
	if err != nil {
		log.Fatalf("could not decrypt balance: %v", err)
		return nil, err
	}
	//fmt.Println("=====Function GetSalary:", account, salary)
	return salary, nil
}

//PaySalary pays salary to a given addresss .
func PaySalary(sim *backends.SimulatedBackend, auth *bind.TransactOpts, contract *Contract, account string) (result bool, err error) {
	//fmt.Println("=====Function PaySalary:", account)
	_, err = contract.PaySalary(auth, common.HexToAddress(account))
	if err != nil {
		log.Fatalf("could not pay salary: %v", err)
		return false, err
	}
	sim.Commit()
	return true, nil
}

//TakeMoney withdrawals money from a account's balance
func TakeMoney(sim *backends.SimulatedBackend, auth *bind.TransactOpts, contract *Contract, account string, amount *big.Int, priv *myCrypto.PrivateKey) (result bool, err error) {
	//fmt.Println("=====Function TakeMoney:", account, amount)
	balanceEncrypted, err := contract.CheckBalance(nil, common.HexToAddress(account))
	if err != nil {
		log.Fatalf("could not get balance: %v", err)
		return false, err
	}
	//!!!!here should make the mod operation in myCrypto.go
	C := balanceEncrypted.C
	D := new(big.Int).Exp(priv.G, amount, priv.P)
	D.ModInverse(D, priv.P)
	D.Mul(D, balanceEncrypted.D)
	D.Mod(D, priv.P)
	//fmt.Println("the offline C,D:", C, D)
	proof, cArray, dArray, err := myCrypto.GenerateRangeProof(priv, C, D)
	if err != nil {
		//log.Fatalf("could not generate range proof: %v", err)
		return false, errors.New("can not generate range proof")
	}
	_, err = contract.TakeMoney(auth, common.HexToAddress(account), amount, proof, cArray, dArray)
	if err != nil {
		log.Fatalf("could not take money: %v", err)
		return false, err
	}
	sim.Commit()

	return true, nil
}

//InitWithSim with simulated env, TBA
func InitWithSim() (sim *backends.SimulatedBackend, auth *bind.TransactOpts, sc *Contract, err error) {

	var keyTest = `
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

	//Create a genesis account with a given keystore here, only for test
	auth, err = bind.NewTransactor(strings.NewReader(keyTest), pwdTest)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
		return nil, nil, nil, err
	}
	fmt.Println("The EOA address:", "0x1a9ec3b0b807464e6d3398a59d6b0a369bf422fa")

	//make the genesis account with a bunch of Ether
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1337000000000)}

	//initiate the simlated backend instead of using testnet, the default gaslimte is 10000000
	//if gas limite < 10000000, can not deploy contract
	//max gaslimit = 9223372036854775807
	sim = backends.NewSimulatedBackend(alloc, 10000000)

	//deploy the contract which is metaCoin.go generated by abigen
	fmt.Println("======Deploy the contract...")
	addr, txDp, sc, err := DeployContract(auth, sim)

	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
		return nil, nil, nil, err
	}
	fmt.Println("The deploy tx hash:", "0x"+common.Bytes2Hex(txDp.Hash().Bytes()))
	fmt.Println("Contract pending deploy:", "0x"+common.Bytes2Hex(addr.Bytes()))

	fmt.Println("=====Commit all pending txs and fresh state.")
	sim.Commit()

	//startTime := time.Now()
	//fmt.Println("Tx start @:", time.Now())
	ctx := context.Background()
	addrDp, err := bind.WaitDeployed(ctx, sim, txDp)
	if err != nil {
		log.Fatalf("failed to deploy contact when mining :%v", err)
		return nil, nil, nil, err
	}
	//check if the contract is deployed
	//fmt.Printf("tx mining take time:%s\n", time.Now().Sub(startTime))
	if bytes.Compare(addr.Bytes(), addrDp.Bytes()) != 0 {
		log.Fatalf("mined address :%s,before mined address:%s", addrDp, addr)
		return nil, nil, nil, err
	}

	return
}

const keyString = "070aad7ed2399348451b3a92091383f4082a780534c30ad59f7d72bc2971e0ab"

//Init with a given environment
func Init() {
	//string priv_key = ""
	//key privpriv_key :=
	keyPrivByte := []byte(keyString)
	fmt.Println(len(keyPrivByte))
	//bind.NewTransactor()
	//bind.NewKeyedTransactor()
	privKey, err := crypto.ToECDSA(keyPrivByte)
	if err != nil {
		log.Fatalf("failed to get the private key:%v", err)
	}
	fmt.Println(privKey.PublicKey)
	//pubBytes := crypto.FromECDSAPub(&priv_key.Public())
	//addr_bytes := crypto.Keccak256(pubBytes[1:])[12:]
	//addr_user := "0x" + common.Bytes2Hex(addr_bytes)
	//fmt.Println("The EOA address:", addr_user)
	//auth := bind.NewKeyedTransactor(priv_key) //for test, use a given key here
}
