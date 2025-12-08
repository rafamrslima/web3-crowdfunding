// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {ERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract MockUSDC is ERC20, ERC20Permit, Ownable {
    uint8 private constant _DECIMALS = 6;

    constructor() 
        ERC20("USD Coin", "USDC") 
        ERC20Permit("USD Coin")
        Ownable(msg.sender)
    {
        _mint(msg.sender, 1_000_000 * 10**_DECIMALS);
    }

    function decimals() public pure override returns (uint8) {
        return _DECIMALS;
    }

    function mint(address to, uint256 amount) external onlyOwner {
        _mint(to, amount);
    }
}