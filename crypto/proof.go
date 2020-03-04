package crypto

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"salary_demo/utils"
)

// GenerateRangeProof makes a range proof
func GenerateRangeProof(priv *PrivateKey, C, D *big.Int) (proof, cArray, dArray []*big.Int, err error) {
	//fmt.Println("+++++ Here is generate range proof")
	s, err := DecryptInt(priv, C, D)
	if err != nil {
		return
	}

	sArray := utils.ConvertToBinArray(s)
	proof = make([]*big.Int, 32*3) //proof is a 32*3 array
	cArray = make([]*big.Int, 32)
	dArray = make([]*big.Int, 32)

	var i int64
	c31 := big.NewInt(1)
	for i = 0; i <= 30; i++ {
		ci, _ := rand.Int(rand.Reader, priv.P) //ci=ranom() should be less than P?
		cArray[i] = ci
		c31.Mul(c31, ci)                                                      //to calculate C_1C_2...C_30
		c31.Mod(c31, priv.P)                                                  //should mod P here?
		cix := new(big.Int).Exp(ci, priv.X, priv.P)                           //to calculate ci^x
		tmpI := new(big.Int).Exp(big.NewInt(2), big.NewInt(i), big.NewInt(0)) //to calculate 2^i
		mk := new(big.Int).Exp(big.NewInt(2), big.NewInt(i), big.NewInt(0))   //a new copy of 2^i, which is the input of SetMembership

		tmpI.Mul(tmpI, sArray[i]) //to calculate 2^i*si
		//there is no need to mod p here since si= 0 or 1
		tmp := new(big.Int).Exp(priv.G, tmpI, priv.P) //tmp=g^(2^i*si)
		dArray[i] = new(big.Int).Mul(cix, tmp)        //di=ci^x * g^(2^i*si)
		dArray[i] = dArray[i].Mod(dArray[i], priv.P)  //do not forget to mod again!
		proof[i*3], proof[i*3+1], proof[i*3+2], err = SetMembershipProof(priv, cArray[i], dArray[i], mk, sArray[i])
		//VerifyMembership(priv, cArray[i], dArray[i], mk, proof[i*3], proof[i*3+1], proof[i*3+2])
	}
	cArray[31] = c31.ModInverse(c31, priv.P) //to calculate 1/(C_1C_2...C_30)
	cArray[31].Mul(C, cArray[31])            //to calculate C/(C_1C_2...C_30)
	cArray[31].Mod(cArray[31], priv.P)       //maybe here need mod operation

	c31x := new(big.Int).Exp(cArray[31], priv.X, priv.P)                    //C31^x
	tmp31 := new(big.Int).Exp(big.NewInt(2), big.NewInt(31), big.NewInt(0)) //2^31
	m31 := new(big.Int).Exp(big.NewInt(2), big.NewInt(31), big.NewInt(0))   //2^31
	tmp31.Mul(tmp31, sArray[31])                                            //2^31*s31
	d31 := new(big.Int).Exp(priv.G, tmp31, priv.P)                          //g^(2^31*s31])
	dArray[31] = c31x.Mul(c31x, d31)                                        //(c31^x)*[g^(2^31*s31])]
	dArray[31].Mod(dArray[31], priv.P)
	proof[31*3], proof[31*3+1], proof[31*3+2], err = SetMembershipProof(priv, cArray[31], dArray[31], m31, sArray[31])
	//VerifyMembership(priv, cArray[31], dArray[31], m31, proof[31*3], proof[31*3+1], proof[31*3+2])
	if err != nil {
		return
	}
	return proof, cArray, dArray, nil
}

//SetMembershipProof returns membership proof.
//Here we input the parameter domain D as single integer m2, since D = {m1=0,m2=2^i} and j=si
func SetMembershipProof(priv *PrivateKey, P, Q, m2, j *big.Int) (e1, s1, s2 *big.Int, err error) {
	m, _ := DecryptInt(priv, P, Q)
	pSubOne := new(big.Int).Sub(priv.P, big.NewInt(1)) //p-1
	v, _ := rand.Int(rand.Reader, pSubOne)             //v=random() and less than P-1
	if m.Cmp(m2) == 0 {
		// suppose m=m2, which means j=2, and the ring starts from 2
		Y2 := new(big.Int).Exp(priv.G, v, priv.P)   //Y2=g^v
		Z2 := new(big.Int).Exp(P, v, priv.P)        //Z2=P^v
		e1 = Hash(Y2, Z2)                           //e1=H(Y2,Z2)
		s1, _ = rand.Int(rand.Reader, pSubOne)      //s1=random() and less than P-1
		tmp := new(big.Int).Exp(priv.H, e1, priv.P) //h^e1
		tmp.ModInverse(tmp, priv.P)                 //h^-e1
		Y1 := new(big.Int).Exp(priv.G, s1, priv.P)  //g^s1
		Y1.Mul(Y1, tmp)                             //Y1=g^s1*h^-e1
		Y1.Mod(Y1, priv.P)                          //mod p
		tmp = new(big.Int).Exp(Q, e1, priv.P)       //Q^e1, because m1=0 so g^m1=1 and Q/g^m1=Q
		tmp.ModInverse(tmp, priv.P)                 //Q^-e1
		Z1 := new(big.Int).Exp(P, s1, priv.P)       //P^s1
		Z1.Mul(Z1, tmp)                             //P^s1 * Q^-e1
		Z1.Mod(Z1, priv.P)                          //mod p
		e2 := Hash(Y1, Z1)                          //e2=H(Y1,Z1)
		s2 = new(big.Int).Mul(e2, priv.X)           //e2*x
		//s2.Mod(s2, pSubOne)                         //mod p-1
		s2.Add(s2, v) //s2=e2*x+v
		//here do not need more mod operation
		return e1, s1, s2, nil
	}
	//m=m1, which means j=1, and the ring starts from 1
	Y1 := new(big.Int).Exp(priv.G, v, priv.P)   //Y1 = g^v
	Z1 := new(big.Int).Exp(P, v, priv.P)        //Z1=P^v
	e2 := Hash(Y1, Z1)                          //e2=H(Y1,Z1)
	s2, _ = rand.Int(rand.Reader, pSubOne)      //s1=random() and less than P-1
	tmp := new(big.Int).Exp(priv.H, e2, priv.P) //h^e2
	tmp.ModInverse(tmp, priv.P)                 //h^-e2
	Y2 := new(big.Int).Exp(priv.G, s2, priv.P)  //g^s2
	Y2.Mul(Y2, tmp)                             //Y2=g^s2*h^-e2
	Y2.Mod(Y2, priv.P)                          //mod p
	tmp = new(big.Int).Exp(priv.G, m2, priv.P)  //g^m2
	tmp.ModInverse(tmp, priv.P)                 //g^-m2
	tmp.Mul(tmp, Q)                             //Q*g^-m2
	tmp.Mod(tmp, priv.P)                        //mod p
	tmp.Exp(tmp, e2, priv.P)                    //(Q*g^-m2)*e2
	tmp.ModInverse(tmp, priv.P)                 //(Q*g^-m2)*-e2
	Z2 := new(big.Int).Exp(P, s2, priv.P)       //P^s2
	Z2.Mul(Z2, tmp)                             //Z2=P^s2*(Q*g^-m2)*-e2
	Z2.Mod(Z2, priv.P)                          //mod p
	e1 = Hash(Y2, Z2)                           //e1=H(Y2,Z2)
	s1 = new(big.Int).Mul(e1, priv.X)           //e1*x
	//s1.Mod(s1, pSubOne)                         //mod p-1
	s1.Add(s1, v) //s1=e1*x+v
	return e1, s1, s2, nil
}

//VerifyMembership is the local verifiy function
func VerifyMembership(priv *PrivateKey, P, Q, mk, proof1, proof2, proof3 *big.Int) bool {
	tmp := new(big.Int).Exp(priv.H, proof1, priv.P) //h^e1
	tmp.ModInverse(tmp, priv.P)                     //h^-e1
	Y1 := new(big.Int).Exp(priv.G, proof2, priv.P)  //g^s1
	Y1 = Y1.Mul(Y1, tmp)                            //Y1=g^s1 * h^-e1
	Y1.Mod(Y1, priv.P)                              //mod p
	tmp = new(big.Int).Exp(Q, proof1, priv.P)       //Q^e1
	fmt.Println("Q^e1:", tmp)
	tmp.ModInverse(tmp, priv.P) //Q^-e1
	fmt.Println("Q^-e1:", tmp)
	Z1 := new(big.Int).Exp(P, proof2, priv.P) //P^s1
	fmt.Println("P^s1:", Z1)
	Z1 = Z1.Mul(Z1, tmp) //P^s1 * Q^-e1
	Z1.Mod(Z1, priv.P)   //mod p
	fmt.Println("P^s1 * Q^-e1", Z1)
	e2 := Hash(Y1, Z1)

	tmp = new(big.Int).Exp(priv.H, e2, priv.P)
	tmp.ModInverse(tmp, priv.P)
	Y2 := new(big.Int).Exp(priv.G, proof3, priv.P)
	Y2 = Y2.Mul(Y2, tmp)
	Y2.Mod(Y2, priv.P)
	tmp = new(big.Int).Exp(priv.G, mk, priv.P)
	tmp.ModInverse(tmp, priv.P)
	tmp = tmp.Mul(tmp, Q)
	tmp.Mod(tmp, priv.P)
	tmp = tmp.Exp(tmp, e2, priv.P)
	tmp.ModInverse(tmp, priv.P)
	Z2 := new(big.Int).Exp(P, proof3, priv.P)
	Z2 = Z2.Mul(Z2, tmp)
	Z2.Mod(Z2, priv.P)
	e1 := Hash(Y2, Z2)
	if e1.Cmp(proof1) == 0 {
		//fmt.Println("Verify Ci:", P, "Di:", Q, "m2=2^i", mk)
		//fmt.Println("e1,s1,s2:",proof,s1,sk)
		fmt.Println("H,P,Q,mk,e1,s1,sk:", priv.H, P, Q, mk, proof1, proof2, proof3)
		fmt.Println("e1,s1,s2:", proof1, proof2, proof3)
		fmt.Println("Y1,Z1,Y2,Z2:", Y1, Z1, Y2, Z2)
		fmt.Println("Verfify true!\n")
		return true
	}
	//fmt.Println("Verify Ci:", P, "Di:", Q, "m2=2^i", mk)
	fmt.Println("e1,s1,s2:", proof1, proof2, proof3)
	fmt.Println("Y1,Z1,Y2,Z2:", Y1, Z1, Y2, Z2)
	//fmt.Println("Proof:", e1, s1, s2)
	fmt.Println("Verfify false!\n")
	return false
}
