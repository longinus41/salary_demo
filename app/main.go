package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"salary_demo/command"

	"github.com/urfave/cli"
)

const generator = 7
const prime = 150001
const privateKey = 113

//Env represents the environment
var Env *command.Environment

//Account represents the account information
var Account *command.Account

func start() {
	fmt.Println("start command")
	var (
		op    string
		parm1 string
		parm2 string
		Input string
	)

	f := bufio.NewReader(os.Stdin)
L:
	for {
		fmt.Print(">")
		Input, _ = f.ReadString('\n')
		if len(Input) == 1 {
			continue
		}
		fmt.Sscan(Input, &op, &parm1, &parm2)
		switch op {
		case "init":
			if parm1 == "" && parm2 == "" {
				Env, Account, _ = command.InitSim()
			} else {
				Env, Account, _ = command.Init(parm1, parm2)
			}
		case "setbalance":
			command.SetBalance(Env, Account, parm1, parm2)
		case "getbalance":
			command.GetBalance(Env, Account, parm1)
		case "setsalary":
			command.SetSalary(Env, Account, parm1, parm2)
		case "getsalary":
			command.GetSalary(Env, Account, parm1)
		case "paysalary":
			command.PaySalary(Env, Account, parm1)
		case "withdraw":
			command.Withdraw(Env, Account, parm1, parm2)
		case "test":
			command.Test()
		case "stop":
			break L
		default:
			fmt.Println("invalid command")
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "salary"
	app.Version = "0.0.1"
	app.Usage = "This is a demo of decrypted salary system"
	app.Action = func(c *cli.Context) error {
		start()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
