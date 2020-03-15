package command

import (
	"fmt"
	"log"
	"math/big"
	"salary_demo/contract"
)

//Test is a test function
//func Test(env *Environment, account *Account) bool {
//	if env == nil || account == nil {
//		fmt.Println("init the envrionment first!")
//		return false
//	}
//	fmt.Println("test...")
//	//contract.HashTest(env.Backend, account.Auth, env.Contract)
//	myConfig.Init()
//	return true
//}

//Test TBA.
func Test() {
	address := "0x1a9ec3b0b807464e6d3398a59d6b0a369bf422fa"
	env, account, _ := InitSim()
	result, _ := contract.GetBalance(env.Contract, address, account.AccountKey)
	if result.Cmp(big.NewInt(0)) != 0 {
		log.Fatalf("Test: the init balance should be 0")
	}
	balance := big.NewInt(1000)
	contract.SetBalance(env.Backend, account.Auth, env.Contract, address, balance, account.AccountKey)
	result, _ = contract.GetBalance(env.Contract, address, account.AccountKey)
	if result.Cmp(balance) != 0 {
		log.Fatalf("Test: the balance set failed")
	}
	salary := big.NewInt(500)
	contract.SetSalary(env.Backend, account.Auth, env.Contract, address, salary, account.AccountKey)
	result, _ = contract.GetSalary(env.Contract, address, account.AccountKey)
	if result.Cmp(salary) != 0 {
		log.Fatalf("Test: the salary set failed")
	}

	contract.PaySalary(env.Backend, account.Auth, env.Contract, address)
	result, _ = contract.GetBalance(env.Contract, address, account.AccountKey)
	if result.Cmp(balance.Add(balance, salary)) != 0 {
		log.Fatalf("Test: pay salary failed")
	}
	fmt.Println("Pay salary success")

	withdraw := big.NewInt(900)
	contract.TakeMoney(env.Backend, account.Auth, env.Contract, address, withdraw, account.AccountKey)
	result, _ = contract.GetBalance(env.Contract, address, account.AccountKey)
	if result.Cmp(balance.Sub(balance, withdraw)) != 0 {
		log.Fatalf("Test: withdraw failed")
	}
	fmt.Println("Withdraw success")
}
