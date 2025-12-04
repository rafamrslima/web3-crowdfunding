// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script} from "forge-std/Script.sol";
import {CrowdFunding} from "../src/CrowdFunding.sol";

contract CrowdFundingScript is Script {
    CrowdFunding public crowdFunding;

    function setUp() public {}

    function run() public {
        vm.startBroadcast();
        address usdc = vm.envAddress("USDC_ADDRESS");

        crowdFunding = new CrowdFunding(usdc);

        vm.stopBroadcast();
    }
}
