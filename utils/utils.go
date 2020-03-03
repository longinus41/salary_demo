package utils

import (
	"errors"
	"math/big"
	"unsafe"
)

const Red = "\033[1;31m%s\033[0m\n"
const Blue = "\033[1;36m%s\033[0m\n"
const Green = "\033[1;32m%s\033[0m\n"

//M is the upper bound of brute force algorithm.
const M = 100000

//Log uses the brute force algorithm for logarithmic computation.
func Log(x, y, p *big.Int) (z *big.Int, err error) {
	var i int64
	for i = 0; i <= M; i++ {
		iBig := big.NewInt(i)
		if new(big.Int).Exp(x, iBig, p).Cmp(y) == 0 {
			z = iBig
			return z, nil
		}
	}
	return nil, errors.New("Error: cannot find discrete logarithm")
}

//FromHex trans a hex number in string to a big integer
func FromHex(hex string) *big.Int {
	n, ok := new(big.Int).SetString(hex, 16)
	if !ok {
		panic("failed to parse hex number")
	}
	return n
}

//ToInt32 converts int64 number to int32 number.
func ToInt32(long int64) int {
	return *(*int)(unsafe.Pointer(&long))
}

//ConvertToBinArray converts decimal number to a binary number in a ascending sort array.
//e.g. ConvertToBinArray(300) = [0 0 1 1 0 1 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
func ConvertToBinArray(n *big.Int) [](*big.Int) {
	nInt32 := ToInt32(n.Int64()) //transfer to int32
	array := make([]*big.Int, 32)
	for i := 31; i >= 0; i-- {
		//fmt.Print(n >> i & 1)
		binBit := nInt32 >> i & 1
		array[i] = big.NewInt(int64(binBit))
	}
	return array
}
