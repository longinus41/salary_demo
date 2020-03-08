pragma solidity >=0.4.25 <0.6.2;


// This is just a simple example of a salary-like contract.
contract Salary {
	struct encrpyt_data{
		uint256 c;
		uint256 d;
	}

	mapping (address => encrpyt_data) public balances;
	mapping (address => encrpyt_data) public salary;
	mapping (address => uint256) public publicKey;

	uint256 public generator;
	uint256 public prime;

	event SetBalance(address indexed _to, uint256 _value1, uint256 _value2);
	event SetSalary(address indexed _to, uint256 _value1, uint256 _value2);
	event PaySalary(address indexed _to);
	event TakeMoney(address indexed _to, uint256 _value);

	constructor() public {
		balances[msg.sender].c = 0;
		balances[msg.sender].d = 0;
		salary[msg.sender].c = 0;
		salary[msg.sender].d = 0;
		//For test here.
		generator = 7;
		prime = 150001;
	}

	function checkBalance(address _addr) public view returns(uint256 _c,uint256 _d) {
		//need some require such as publickey!=null
		return (balances[_addr].c, balances[_addr].d);
	}

	function setBalance(address _addr,uint256 _c,uint256 _d,uint256 _h) public returns (bool){
		balances[_addr].c = _c;
		balances[_addr].d = _d;
		publicKey[_addr] = _h;	//Here we update the publickey in the set function
		emit SetBalance(_addr,_c,_d);
		return true;
	}

	function checkSalary(address _addr) public view returns(uint256 _c,uint256 _d) {
		//need some require such as publickey!=null
		return (salary[_addr].c, salary[_addr].d);
	}

	function setSalary(address _addr,uint256 _c,uint256 _d) public returns (bool){
		salary[_addr].c = _c;
		salary[_addr].d = _d;
		emit SetSalary(_addr,_c,_d);
		return true;
	}

	function paySalary(address _addr) public returns (bool){
		//need some require such as publickey!=null
		balances[_addr].c = mulmod(balances[_addr].c,salary[_addr].c,prime);
		balances[_addr].d = mulmod(balances[_addr].d,salary[_addr].d,prime);
		emit PaySalary(_addr);
		return true;
	}

	function takeMoney(address _addr,uint256 _t, uint256[] memory _proof, uint256[] memory _cArray, uint256[] memory _dArray)
	public
	returns (bool){
		require(_t>0,'Money should be a positive.');
		uint256 tmpC = balances[_addr].c;
		uint256 tmpD = modInverse(modExp(generator,_t,prime), prime);
		tmpD = mulmod(balances[_addr].d, tmpD, prime);
		if (verifyRangeProof(_addr, _proof, _cArray, _dArray,tmpC, tmpD)) {
			balances[_addr].d = tmpD;
			//remit money here.
			emit TakeMoney(_addr, _t);
			return true;
		}else
			return false;
	}


	function verifyRangeProof(address _addr, uint256[] memory _proof, uint256[] memory _cArray, uint256[] memory _dArray, uint256 _c, uint256 _d)
	public view
	returns (bool result){
		uint256 productC = 1;
		uint256 productD = 1;
		result = true;
		for(uint i = 0; i <= 31; i += 1){
			productC = mulmod(productC,_cArray[i],prime);
			productD = mulmod(productD,_dArray[i],prime);
			result = result && verifyMembership(_addr,_cArray[i],_dArray[i],(2**i),_proof[i*3],_proof[i*3+1],_proof[i*3+2]);
		}
		result = result && ((_c == productC) ? true : false);
		result = result && ((_d == productD) ? true : false);
		return result;
	}

	function verifyMembership(address _addr, uint256 _p, uint256 _q, uint _mk,uint256 _proof1, uint256 _proof2,uint256 _proof3)
	public view
	returns (bool){
		//proof={proof1,proof2,proof3}={e1,s1,sk}
		//domain={m1,mk}={0,2^i}
		uint256 tmp = modInverse(modExp(publicKey[_addr],_proof1,prime),prime);//h^-e1
		uint256 Y1 = modExp(generator,_proof2,prime);						//g^s1
		Y1 = mulmod(Y1, tmp, prime);										//Y1 = g^s1 * h^-e1
		uint256 Z1 = modExp(_p,_proof2,prime);							    //P^s1
		tmp = modExp(_q,_proof1, prime);									//[Q/(g^m1)]^e1
		tmp = modInverse(tmp,prime);										//[Q/(g^m1)]^-e1
		Z1 = mulmod(Z1, tmp, prime);										//Z1 = P^s1 * [Q/(g^m1)]^-e1

		uint256 e2 = simpleHash(Y1,Z1);
		tmp = modInverse(modExp(publicKey[_addr],e2,prime),prime);          //h^-ek, where ek=e2
		uint256 Yk = modExp(generator,_proof3,prime);						//g^sk
		Yk = mulmod(Yk, tmp, prime);										//Yk = g^sk * h^-ek
		uint256 Zk = modExp(_p,_proof3,prime);							    //P^sk
		tmp = modExp(generator,_mk,prime);									//g^mk
		tmp = modInverse(tmp, prime);										//1/(g^mk)
		tmp = mulmod(_q, tmp, prime);										//Q/(g^mk)
		tmp = modExp(tmp, e2, prime);									    //[Q/(g^mk)]^ek
		tmp = modInverse(tmp,prime);										//[Q/(g^mk)]^-ek
		Zk = mulmod(Zk, tmp, prime);										//Zk = P^sk * [Q/(g^mk)]^-ek

		return (_proof1 == simpleHash(Yk,Zk)?true:false);
	}

	//Calculate x**y mod |m| using a brute force approach, more efficient method will be added in the future.
	function modExp(uint256 _base, uint256 _exp, uint _mod) internal pure returns (uint256 result){
		require(_base>0 && _exp>=0 && _mod>0,'x,y,z should be positives.');
		if(_exp==0){
			return 1;
		}
		result = 1;
		uint256 _tmpBase = _base;
		uint256 _tmpExp = _exp;
		uint256 _tmpMod = _mod;
		for(uint count = 1; count <= _tmpExp; count *= 2){
			if(_tmpExp & count != 0)
				result = mulmod(result, _tmpBase, _tmpMod);
			_tmpBase = mulmod(_tmpBase,_tmpBase,_tmpMod);
		}
		return result;
	}

	//A more gas friendly method of modular exponentiation using precompiled contract in Ethereum. Only on Ethereum full node.
	function modExpQuick(uint256 _b, uint256 _e, uint256 _m) public returns (uint256 result) {
        assembly {
            // Free memory pointer
            let pointer := mload(0x40)

            // Define length of base, exponent and modulus. 0x20 == 32 bytes
            mstore(pointer, 0x20)
            mstore(add(pointer, 0x20), 0x20)
            mstore(add(pointer, 0x40), 0x20)

            // Define variables base, exponent and modulus
            mstore(add(pointer, 0x60), _b)
            mstore(add(pointer, 0x80), _e)
            mstore(add(pointer, 0xa0), _m)

            // Store the result
            let value := mload(0xc0)

            // Call the precompiled contract 0x05 = bigModExp
            if iszero(call(not(0), 0x05, 0, pointer, 0xc0, value, 0x20)) {
                revert(0, 0)
            }
            result := mload(value)
        }
    }

	//Calculate the modular multiplicative invserse of x using Fermat's Little Theorem
	function modInverse(uint256 _g, uint256 _n) internal pure returns (uint256){
		return modExp(_g, _n-2, _n);
	}

	function simpleHash(uint256 _x, uint256 _y) public pure returns(uint256) {
		return _x+_y;
	}

	function myHash(uint256 _x, uint256 _y) public pure returns(uint256){
		uint256 abs;
		if(_x>_y){
			abs = _x-_y;
		}else{
			abs = _y-_x;
		}
		bytes32 b = keccak256(abi.encode(abs));
		return bytesToUint(b);
	}


	function bytesToUint(bytes32 b) public pure returns (uint256){
		uint256 number;
		for(uint i = 0; i < b.length; i++){
            number = number + uint8(b[i])*(2**(8*(b.length-(i+1))));
        }
        return number;
	}

}
