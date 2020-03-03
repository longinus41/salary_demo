package crypto

import (
	"math/big"
)

//Hash returns a hash of x,y TBA
func Hash(X, Y *big.Int) (h *big.Int) {
	//sha3 := sha3.New256()
	//sha3.Write(X.Bytes())
	//res := sha3.Sum(Y.Bytes())
	//result := hex.EncodeToString(res)
	//h = new(big.Int).SetBytes(res)
	//fmt.Println("Hash result:", res, h)
	//return X.Add(X, Y)
	h = new(big.Int).Add(X, Y)
	//h.Mod(h, priv.v)
	return
}
