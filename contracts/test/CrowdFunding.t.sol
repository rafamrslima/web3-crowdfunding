// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test} from "forge-std/Test.sol";
import {CrowdFunding} from "../src/CrowdFunding.sol";
import {MockUSDC} from "../src/MockUSDC.sol";

contract CrowdFundingTest is Test {
    CrowdFunding public crowdFunding;
    MockUSDC public usdc;
    address public campaignOwner;
    address public donor;

    function setUp() public {
        usdc = new MockUSDC();
        crowdFunding = new CrowdFunding(address(usdc));
        campaignOwner = makeAddr("campaignOwner");
        donor = makeAddr("donor");
        
        // Mint USDC to donor for testing
        usdc.mint(donor, 1000 * 1e6); // 1000 USDC
    }

    function test_CreateCampaign() public {
        vm.prank(campaignOwner);
        uint256 campaignId = crowdFunding.createCampaign(
            5 * 1e6, // 5 USDC
            block.timestamp + 30 days,
            keccak256(abi.encodePacked("hash"))
        );
        
        assertEq(campaignId, 0);
        assertEq(crowdFunding.numberOfCampaigns(), 1);
    }

    function test_RefundDonor_Success() public {
        // Create a campaign that won't reach its target
        vm.prank(campaignOwner);
        uint256 campaignId = crowdFunding.createCampaign(
            1000 * 1e6, // 1000 USDC target (high target to ensure goal not reached)
            block.timestamp + 1 days, // Short deadline
            keccak256(abi.encodePacked("refund_test"))
        );

        // Donor makes a donation
        uint256 donationAmount = 100 * 1e6; // 100 USDC
        vm.startPrank(donor);
        usdc.approve(address(crowdFunding), donationAmount);
        crowdFunding.donateToCampaign(campaignId, donationAmount);
        vm.stopPrank();

        // Fast forward past the campaign deadline
        vm.warp(block.timestamp + 2 days);

        // Check balances before refund
        uint256 donorBalanceBefore = usdc.balanceOf(donor);
        uint256 contractBalanceBefore = usdc.balanceOf(address(crowdFunding));

        // Donor requests refund
        vm.prank(donor);
        crowdFunding.refundDonor(campaignId);

        // Verify refund was successful
        uint256 donorBalanceAfter = usdc.balanceOf(donor);
        uint256 contractBalanceAfter = usdc.balanceOf(address(crowdFunding));

        assertEq(donorBalanceAfter, donorBalanceBefore + donationAmount, "Donor should receive full refund");
        assertEq(contractBalanceAfter, contractBalanceBefore - donationAmount, "Contract balance should decrease by refund amount");
    }
}