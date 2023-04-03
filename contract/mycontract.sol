// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;

contract MyContract {
    string public message;

    constructor() {
        message = "Hello, world!";
    }

    function setMessage(string memory newMessage) public {
        message = newMessage;
    }

    function getMessage() public view returns (string memory) {
        return message;
    }
}
