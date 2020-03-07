package crypto

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"golang.org/x/crypto/sha3"
)

//Hash returns a hash of x,y TBA
func simpleHash(X, Y *big.Int) (h *big.Int) {
	h = new(big.Int).Add(X, Y)
	return
}

//Sha3 returns a sha3-256 of integer X and Y
func Sha3(x, y *big.Int) *big.Int {
	mySHA := sha3.New256()
	mySHA.Write(x.Bytes())
	result := mySHA.Sum(y.Bytes())
	return new(big.Int).SetBytes(result)
}

//Sha3Sol returns a sha3-256 of integer X
func Sha3Sol(x, y *big.Int) *big.Int {
	abs := new(big.Int).Sub(x, y)
	abs.Abs(abs)
	h := sha3.NewLegacyKeccak256()
	h.Write(abi.U256(abs))
	hashBytes := h.Sum(nil)
	return new(big.Int).SetBytes(hashBytes)
}
