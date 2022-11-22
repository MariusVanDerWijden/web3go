package contracts

// Download solc
// wget -nc "https://github.com/ethereum/solidity/releases/download/v0.8.14/solc-static-linux"
// Make executable
// chmod +x solc-static-linux

//go:generate ./solc-static-linux --overwrite --bin --abi -o . CoolContract.sol
//go:generate abigen --pkg contracts --type CoolContract --abi CoolContract.abi --bin CoolContract.bin --out coolContract.go
