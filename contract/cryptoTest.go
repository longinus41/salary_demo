package contract

import (
	"fmt"
	"math/big"
	"salary_demo/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
)

//VerifyMembershipTest is the test func
func VerifyMembershipTest() {
	//H := big.NewInt(66436)
	//P := big.NewInt(23222)
	//Q := big.NewInt(39282)
	//mk := big.NewInt(1)
	//proof1 := big.NewInt(82098)
	//proof2 := big.NewInt(9393474)
	//proof3 := big.NewInt(125120)

	//out, _ := contract.VerifyMembershipTest(nil, H, P, Q, mk, proof1, proof2, proof3)
	//fmt.Println("")
	//fmt.Println("============================")
	//fmt.Println("hash:", out)
}

//HashTest tests the Hash function in go and solidity
func HashTest(sim *backends.SimulatedBackend, auth *bind.TransactOpts, contract *Contract) {
	fmt.Println("=====Hash Test")
	x := big.NewInt(11)
	y := big.NewInt(222)
	hOff := crypto.Sha3Sol(x, y)
	fmt.Println("sha256 offchain:", hOff)
	//fmt.Println("sha256 offchain:", new(big.Int).SetBytes(hOff))

	hOn, _ := contract.MyHash(nil, x, y)
	//fmt.Println("bytes onchain:", b.H1)
	fmt.Println("keccak256 onchain:", hOn)
	//fmt.Println("keccak256 onchain:", b)

}
