#!  /bin/bash

fileName="Salary.sol"
abiFileName="Salary_sol_Salary.abi"
binFileName="Salary_sol_Salary.bin"
goFileName="salary.go"

currentPath=$(pwd)
echo "当前文件夹路径: $currentPath"

echo "start to build abi"
solcjs $fileName -o $currentPath --abi

echo "start to build bin"
solcjs $fileName -o $currentPath --bin

echo "start to transfer to go api"
abigen --abi $abiFileName --bin $binFileName --pkg contract --out $currentPath/$goFileName