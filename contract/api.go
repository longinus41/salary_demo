package contract

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	myCrypto "salary_demo/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
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
	if sim != nil {
		sim.Commit()
	}
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
	if sim != nil {
		sim.Commit()
	}
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
	if sim != nil {
		sim.Commit()
	}

	return true, nil
}

//CalHash returns keccak256 hash of two integers
//func CalHash(contract *Contract, x, y *big.Int) (hash *big.Int, err error) {
//	hash, err = contract.Hash(nil, x, y)
//	if err != nil {
//		return nil, errors.New("can not generate hash")
//	}
//	return hash, nil
//}
