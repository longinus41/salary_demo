// Package crypto implements ElGamal encryption, suitable for OpenPGP,
// as specified in "A Public-Key Cryptosystem and a Signature Scheme Based on
// Discrete Logarithms," IEEE Transactions on Information Theory, v. IT-31,
// n. 4, 1985, pp. 469-472.
//
// This form of ElGamal embeds PKCS#1 v1.5 padding, which may make it
// unsuitable for other protocols. RSA should be used in preference in any
// case.
package crypto

import (
	"crypto/rand"
	"errors"
	"io"
	"math/big"
	"mycrypt/utils"
)

// PublicKey represents an ElGamal public key.
type PublicKey struct {
	G, P, H *big.Int //H:= G^X mod P
}

// PrivateKey represents an ElGamal private key.
type PrivateKey struct {
	PublicKey
	X *big.Int
}

// EncryptInt encrypts the given integers with the given public key.
func EncryptInt(random io.Reader, pub *PublicKey, m *big.Int) (C, D *big.Int, err error) {

	y, err := rand.Int(random, pub.P)
	if err != nil {
		return
	}

	C = new(big.Int).Exp(pub.G, y, pub.P)
	tmp := new(big.Int).Exp(pub.H, y, pub.P)

	D = tmp.Mul(tmp, new(big.Int).Exp(pub.G, m, pub.P))
	D.Mod(D, pub.P)

	return
}

// DecryptInt takes two integers, resulting from an ElGamal encryption.
func DecryptInt(priv *PrivateKey, C, D *big.Int) (m *big.Int, err error) {
	tmp := new(big.Int).Exp(C, priv.X, priv.P)
	if tmp.ModInverse(tmp, priv.P) == nil {
		return nil, errors.New("elgamal: invalid private key")
	}
	tmp.Mul(tmp, D)
	tmp.Mod(tmp, priv.P)
	//fmt.Println("func Dec:", tmp)
	return utils.Log(priv.G, tmp, priv.P)
}
