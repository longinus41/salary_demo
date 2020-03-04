package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"salary_demo/command"
	"salary_demo/contract"
	"salary_demo/crypto"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"
)

// This is the 1024-bit MODP group from RFC 5114, section 2.1:
//const primeHex = "B10B8F96A080E01DDE92DE5EAE5D54EC52C99FBCFB06A3C69A6A9DCA52D23B616073E28675A23D189838EF1E2EE652C013ECB4AEA906112324975C3CD49B83BFACCBDD7D90C4BD7098488E9C219A73724EFFD6FAE5644738FAA31A4FF55BCCC0A151AF5F0DC8B4BD45BF37DF365C1A65E68CFDA76D4DA708DF1FB2BC2E4A4371"
//const generatorHex = "A4D1CBD5C3FD34126765A442EFB99905F8104DD258AC507FD6406CFF14266D31266FEA1E5C41564B777E690F5504F213160217B4B01B886A5E91547F9E2749F4D7FBD7D3B9A92EE1909D0D2263F80A76A6A24C087A091F531DBF0A0169B6A28AD662A4D18E73AFA32D779D5918D08BC8858F4DCEF97C2A24855E6EEB22B3B2E5"

const generator = 7
const prime = 150001
const privateKey = 113

func cryptTest() {
	priv := &crypto.PrivateKey{
		PublicKey: crypto.PublicKey{
			//here we use a small generator g=7, a prime num p=150001, and a private key x=113
			G: new(big.Int).SetInt64(generator),
			P: new(big.Int).SetInt64(prime),
		},
		X: new(big.Int).SetInt64(privateKey),
	}
	priv.H = new(big.Int).Exp(priv.G, priv.X, priv.P)

	fmt.Println("+++++The public key:", priv.PublicKey)

	sim, auth, sc, err := contract.InitWithSim()
	if err != nil {
		fmt.Println(err)
		return
	}
	addrTest := "0x" + common.Bytes2Hex(auth.From.Bytes())

	balance, _ := contract.GetBalance(sc, addrTest, priv)
	fmt.Println("+++++The init balance:", balance)

	//set balance=500
	contract.SetBalance(sim, auth, sc, addrTest, big.NewInt(500), priv)
	//so balance=500
	balance, _ = contract.GetBalance(sc, addrTest, priv)

	//set salary=200
	contract.SetSalary(sim, auth, sc, addrTest, big.NewInt(200), priv)
	//pay salary
	contract.PaySalary(sim, auth, sc, addrTest)
	//so balance = 500+200=700
	balance, _ = contract.GetBalance(sc, addrTest, priv)

	//take 100 away
	contract.TakeMoney(sim, auth, sc, addrTest, big.NewInt(100), priv)
	//so balance = 700-100=600
	balance, _ = contract.GetBalance(sc, addrTest, priv)

	//take 200 away
	contract.TakeMoney(sim, auth, sc, addrTest, big.NewInt(700), priv)
	//so balance = 600-200=400
	balance, _ = contract.GetBalance(sc, addrTest, priv)

	//hash test
	//contract.HashTest(sim, auth, sc)

	//VerifyMembershipTest test
	//contract.VerifyMembershipTest(sim, auth, sc)
}

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
			cryptTest()
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
