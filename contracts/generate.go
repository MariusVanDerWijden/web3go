package contracts

// Download solc
// wget -nc "https://github.com/ethereum/solidity/releases/download/v0.8.14/solc-static-linux"
// Make executable
// chmod +x solc-static-linux

//go:generate abigen --pkg contracts --sol CoolContract.sol --out coolContract.go --solc ./solc-static-linux
