# Salary Demo
A simple encrypted salary demo

## Introduction
We design an smart contract for paying salary, in order to achieve the fairness and eradicate unreasonable refusal/delay. Moreover we plan to hide the actual amount of each agent's balance/salary except for herself and trusted party (say, the boss or the board of directors). Since all on-chain data can be analyzed by locally hacking the EVM, the secrets need to be cryptographically encrypted.

There are more details in [https://github.com/freeof123/salary-contract-algo](https://github.com/freeof123/salary-contract-algo).



## Usage

### run

This is a go project. Ensure you have go.

```
$cd app
$go run main.go
```

####initialization

`>init`

It generates genesis account with a given keystore first, and initiate the simulated backend which looks similar to Ethereum testnet. Then the salary contract will be deployed on the backend automatically.

We will provide deployment on Ethereum and testnets like Ganache in the future.

####set balance
`>setbalance [address] [amount]`

It sets the balance of a given account which represented as address to a new value, and public key of the account will be saved in the contract.

####check balance
`>getbalance [address]`

It returns the balance of account.

####set salary
`>setsalar [address] [amount]`

Similar to `setbalance`

####check salary
`>getsalary [address]`

Similar to `setsalary`

####payroll
`>paysalary [address]`

Pay salary to account. For example, the initial balance of account A is _m_, and his salary is _s_, the balance will be _m+s_ after payroll.

####withdraw
`>withdraw [address] [amount]`

Clami a certain amount of money. For example, the initial balance of account A is _m_, and amount withdrawn is _c_, the balance will be _m-c_ after drawings.

## Development

Okay you might need minor code changes in the smart contract for other usage.

In order to be able to compile the program, you need a few things. First and most importantly, you need the _solc_ Solidity compiler.

```
$npm install -g solc
$npm install -g solc-cli
```

Then, just fetch go-ethereum and build it:

```
$go get github.com/ethereum/go-ethereum
$cd $GOPATH/src/github.com/ethereum/go-ethereum/
$make
$make devtools
```

With _solc_ and _geth devtools_ in place, we can start by generating a Go-version of the _Salary.sol_ file, which holds our smart contract:

```
$solcjs Salary.sol -o ./ --abi
$solcjs Salary.sol -o ./ --bin
$abigen --abi Salary_sol_Salary.abi --bin Salary_sol_Salary.bin --pkg contract --out $GOPATH/src/salary_demo/contract/salary.go
```

Or you can just run the _compile.sh_ in /contract.

