// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test} from "forge-std/Test.sol";
import {CrowdFunding} from "../src/CrowdFunding.sol";

contract CrowdFundingTest is Test {
    CrowdFunding public crowdFunding;
    address public campaignOwner;

    function test_CreateCampaign() public {
        campaignOwner = makeAddr("campaignOwner");
        address usdc = vm.envAddress("USDC_ADDRESS");
        crowdFunding = new CrowdFunding(usdc);
        uint256 campaignId = crowdFunding.createCampaign(
            campaignOwner,
            "Test Campaign",
            "A test campaign for crowdfunding",
            5 ether,
            block.timestamp + 30 days,
            "image.jpg"
        );
        
        assertEq(campaignId, 0);
        assertEq(crowdFunding.numberOfCampaigns(), 1);
    }
}