pragma solidity ^0.8.14;
contract CoolContract {
    uint256 balance;
    event Deposited(address addr);
    
    function Deposit() public payable {
        balance += msg.value;
        emit Deposited(msg.sender);
    }
    
    function seeBalance() public view returns (uint256) {
        return balance;
    }
}