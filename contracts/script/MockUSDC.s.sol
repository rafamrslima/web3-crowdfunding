// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {MockUSDC} from "../src/MockUSDC.sol";

contract MockUSDCScript is Script {
    uint256 public mintAmount = 1_000 * 1e6; // 1,000 USDC

    function run() external {
        vm.startBroadcast();
        address recipient = vm.envAddress("RECIPIENT_ADDRESS");

        MockUSDC usdc = new MockUSDC();
        console.log("MockUSDC deployed at:", address(usdc));

        usdc.mint(recipient, mintAmount);
        console.log("Minted USDC to recipient:", recipient);
        console.log("Mint amount (raw):", mintAmount);

        vm.stopBroadcast();
    }
}
