package command

import (
	"fmt"
	"log"
	"math/big"

	//myConfig "salary_demo/config"
	"salary_demo/contract"
	"salary_demo/utils"
)

//SetBalance is the operator for setting balance.
func SetBalance(env *Environment, account *Account, address, balance string) (result bool) {
	if env == nil || account == nil {
		fmt.Println("init the envrionment first!")
		return false
	}
	fmt.Println("set balance...")
	b, err := new(big.Int).SetString(balance, 10)
	if !err {
		log.Fatalf("Error: invalid number")
		return err
	}
	result, _ = contract.SetBalance(env.Backend, account.Auth, env.Contract, address, b, account.AccountKey)
	return result
}

//GetBalance is the operator for getting balance.
func GetBalance(env *Environment, account *Account, address string) bool {
	if env == nil || account == nil {
		fmt.Println("init the envrionment first!")
		return false
	}
	fmt.Println("get balance...")
	result, _ := contract.GetBalance(env.Contract, address, account.AccountKey)
	fmt.Print("balance: ")
	fmt.Printf(utils.Green, result)
	return true
}

//SetSalary is the operator for setting salary.
func SetSalary(env *Environment, account *Account, address, salary string) (result bool) {
	if env == nil || account == nil {
		fmt.Println("init the envrionment first!")
		return false
	}
	fmt.Println("set salary...")
	s, err := new(big.Int).SetString(salary, 10)
	if !err {
		log.Fatalf("Error: invalid number")
		return err
	}
	result, _ = contract.SetSalary(env.Backend, account.Auth, env.Contract, address, s, account.AccountKey)
	return result
}

//GetSalary is the operator for getting salary.
func GetSalary(env *Environment, account *Account, address string) bool {
	if env == nil || account == nil {
		fmt.Println("init the envrionment first!")
		return false
	}
	fmt.Println("get salary...")
	result, _ := contract.GetSalary(env.Contract, address, account.AccountKey)
	fmt.Print("salary: ")
	fmt.Printf(utils.Green, result)
	return true
}

//PaySalary is the operator for payroll.
func PaySalary(env *Environment, account *Account, address string) bool {
	if env == nil || account == nil {
		fmt.Println("init the envrionment first!")
		return false
	}
	fmt.Println("pay salary...")
	result, _ := contract.PaySalary(env.Backend, account.Auth, env.Contract, address)
	return result
}

//Withdraw is the operator for taking money.
func Withdraw(env *Environment, account *Account, address, money string) bool {
	if env == nil || account == nil {
		fmt.Println("init the envrionment first!")
		return false
	}
	fmt.Println("withdraw...")
	m, err := new(big.Int).SetString(money, 10)
	if !err {
		log.Fatalf("Error: invalid number")
		return err
	}
	result, err1 := contract.TakeMoney(env.Backend, account.Auth, env.Contract, address, m, account.AccountKey)
	if err1 != nil {
		log.Println("withdraw failed: ", err1)
		return result
	}
	return result
}
