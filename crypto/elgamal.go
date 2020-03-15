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
	"fmt"
	"io"
	"math/big"
	"salary_demo/utils"
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

// This is the 1024-bit MODP group from RFC 5114, section 2.1:
//const primeHex = "B10B8F96A080E01DDE92DE5EAE5D54EC52C99FBCFB06A3C69A6A9DCA52D23B616073E28675A23D189838EF1E2EE652C013ECB4AEA906112324975C3CD49B83BFACCBDD7D90C4BD7098488E9C219A73724EFFD6FAE5644738FAA31A4FF55BCCC0A151AF5F0DC8B4BD45BF37DF365C1A65E68CFDA76D4DA708DF1FB2BC2E4A4371"
//const generatorHex = "A4D1CBD5C3FD34126765A442EFB99905F8104DD258AC507FD6406CFF14266D31266FEA1E5C41564B777E690F5504F213160217B4B01B886A5E91547F9E2749F4D7FBD7D3B9A92EE1909D0D2263F80A76A6A24C087A091F531DBF0A0169B6A28AD662A4D18E73AFA32D779D5918D08BC8858F4DCEF97C2A24855E6EEB22B3B2E5"

// EncryptInt encrypts the given integers with the given public key.
func EncryptInt(random io.Reader, pub *PublicKey, m *big.Int) (C, D *big.Int, err error) {
	if pub == nil || m == nil {
		return nil, nil, errors.New("elgamal: invalid public key or plaintext")
	}
	y, err := rand.Int(random, pub.P)
	if err != nil {
		return nil, nil, fmt.Errorf("elgamal: random number generate failed: %v", err)
		//errors.New("elgamal: random number generate failed, pub.G=", pub.G)
	}
	C = new(big.Int).Exp(pub.G, y, pub.P)
	tmp := new(big.Int).Exp(pub.H, y, pub.P)
	D = tmp.Mul(tmp, new(big.Int).Exp(pub.G, m, pub.P))
	D.Mod(D, pub.P)

	return
}

// DecryptInt takes two integers, resulting from an ElGamal encryption.
func DecryptInt(priv *PrivateKey, C, D *big.Int) (m *big.Int, err error) {
	if priv == nil || C == nil || D == nil {
		return nil, errors.New("elgamal: invalid private key or ciphertext")
	}
	tmp := new(big.Int).Exp(C, priv.X, priv.P)
	if tmp.ModInverse(tmp, priv.P) == nil {
		return nil, errors.New("elgamal: invalid private key")
	}
	tmp.Mul(tmp, D)
	tmp.Mod(tmp, priv.P)
	return utils.Log(priv.G, tmp, priv.P)
}
