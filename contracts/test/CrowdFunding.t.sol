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
            block.timestamp + 1 days,
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

    function test_DonateToCampaign_Success() public {
        // Create campaign
        vm.prank(campaignOwner);
        uint256 campaignId = crowdFunding.createCampaign(
            100 * 1e6, // 100 USDC target
            block.timestamp + 30 days,
            keccak256(abi.encodePacked("donation_test"))
        );

        uint256 donationAmount = 50 * 1e6; // 50 USDC
        
        // Donor makes donation
        vm.startPrank(donor);
        usdc.approve(address(crowdFunding), donationAmount);
        
        vm.expectEmit(true, true, false, true);
        emit CrowdFunding.DonationReceived(campaignId, donor, donationAmount);
        
        crowdFunding.donateToCampaign(campaignId, donationAmount);
        vm.stopPrank();

        // Verify campaign state updated
        (address _owner, uint256 _target, uint256 _deadline, uint256 amountCollected, bool _withdrawn) 
            = crowdFunding.campaigns(campaignId);
        assertEq(amountCollected, donationAmount);
        assertEq(crowdFunding.contributions(campaignId, donor), donationAmount);
    }

    function test_DonateToCampaign_OwnerCannotDonate() public {
        // Create campaign
        vm.prank(campaignOwner);
        uint256 campaignId = crowdFunding.createCampaign(
            100 * 1e6,
            block.timestamp + 30 days,
            keccak256(abi.encodePacked("owner_donation_test"))
        );

        // Owner tries to donate to own campaign
        usdc.mint(campaignOwner, 100 * 1e6);
        vm.startPrank(campaignOwner);
        usdc.approve(address(crowdFunding), 50 * 1e6);
                
        vm.expectRevert("Owner can't donate");
        crowdFunding.donateToCampaign(campaignId, 50 * 1e6);
        vm.stopPrank();
    }

    function test_Withdraw_Success() public {
        // Create campaign with achievable target
        vm.prank(campaignOwner);
        uint256 campaignId = crowdFunding.createCampaign(
            50 * 1e6, // 50 USDC target
            block.timestamp + 1 days,
            keccak256(abi.encodePacked("withdraw_test"))
        );

        // Donate to reach target
        vm.startPrank(donor);
        usdc.approve(address(crowdFunding), 60 * 1e6);
        crowdFunding.donateToCampaign(campaignId, 60 * 1e6); // Exceed target
        vm.stopPrank();

        // Fast forward past deadline
        vm.warp(block.timestamp + 2 days);

        uint256 ownerBalanceBefore = usdc.balanceOf(campaignOwner);

        // Owner withdraws funds
        vm.prank(campaignOwner);
        vm.expectEmit(true, true, false, true);
        emit CrowdFunding.FundsWithdrawn(campaignId, campaignOwner, 60 * 1e6);
        
        crowdFunding.withdraw(campaignId);

        // Verify withdrawal
        uint256 ownerBalanceAfter = usdc.balanceOf(campaignOwner);
        assertEq(ownerBalanceAfter, ownerBalanceBefore + 60 * 1e6);
        
        (address _owner, uint256 _target, uint256 _deadline, uint256 _amountCollected, bool withdrawn) 
            = crowdFunding.campaigns(campaignId);
        assertTrue(withdrawn);
    }

    function test_Withdraw_FailsIfNotOwner() public {
        vm.prank(campaignOwner);
        uint256 campaignId = crowdFunding.createCampaign(
            50 * 1e6,
            block.timestamp + 1 days,
            keccak256(abi.encodePacked("not_owner_test"))
        );

        // Donate and reach target
        vm.startPrank(donor);
        usdc.approve(address(crowdFunding), 60 * 1e6);
        crowdFunding.donateToCampaign(campaignId, 60 * 1e6);
        vm.stopPrank();

        vm.warp(block.timestamp + 2 days);

        // Non-owner tries to withdraw
        vm.prank(donor);
        vm.expectRevert("Withdraw should be done by the campaign owner");
        crowdFunding.withdraw(campaignId);
    }

    function test_Withdraw_FailsIfTargetNotReached() public {
        vm.prank(campaignOwner);
        uint256 campaignId = crowdFunding.createCampaign(
            100 * 1e6, // 100 USDC target
            block.timestamp + 1 days,
            keccak256(abi.encodePacked("target_not_reached_test"))
        );

        // Donate less than target
        vm.startPrank(donor);
        usdc.approve(address(crowdFunding), 50 * 1e6);
        crowdFunding.donateToCampaign(campaignId, 50 * 1e6); // Below target
        vm.stopPrank();

        vm.warp(block.timestamp + 2 days);

        // Try to withdraw
        vm.prank(campaignOwner);
        vm.expectRevert("Campaign didn't reach the target");
        crowdFunding.withdraw(campaignId);
    }

    function test_GetDonators() public {
        vm.prank(campaignOwner);
        uint256 campaignId = crowdFunding.createCampaign(
            100 * 1e6,
            block.timestamp + 30 days,
            keccak256(abi.encodePacked("get_donators_test"))
        );

        address donor2 = makeAddr("donor2");
        usdc.mint(donor2, 1000 * 1e6);

        // Multiple donations
        vm.startPrank(donor);
        usdc.approve(address(crowdFunding), 30 * 1e6);
        crowdFunding.donateToCampaign(campaignId, 30 * 1e6);
        vm.stopPrank();

        vm.startPrank(donor2);
        usdc.approve(address(crowdFunding), 20 * 1e6);
        crowdFunding.donateToCampaign(campaignId, 20 * 1e6);
        vm.stopPrank();

        // Get donators and donations
        (address[] memory donators, uint256[] memory donations) = crowdFunding.getDonators(campaignId);
        
        assertEq(donators.length, 2);
        assertEq(donations.length, 2);
        assertEq(donators[0], donor);
        assertEq(donators[1], donor2);
        assertEq(donations[0], 30 * 1e6);
        assertEq(donations[1], 20 * 1e6);
    }

    function test_GetCampaigns() public {
        // Create multiple campaigns
        vm.startPrank(campaignOwner);
        crowdFunding.createCampaign(100 * 1e6, block.timestamp + 30 days, keccak256("campaign1"));
        crowdFunding.createCampaign(200 * 1e6, block.timestamp + 60 days, keccak256("campaign2"));
        vm.stopPrank();

        // Get all campaigns
        CrowdFunding.Campaign[] memory campaigns = crowdFunding.getCampaigns();
        
        assertEq(campaigns.length, 2);
        assertEq(campaigns[0].target, 100 * 1e6);
        assertEq(campaigns[1].target, 200 * 1e6);
        assertEq(campaigns[0].owner, campaignOwner);
        assertEq(campaigns[1].owner, campaignOwner);
    }

    function test_CreateCampaign_FailsWithPastDeadline() public {
        vm.prank(campaignOwner);
        vm.expectRevert("The deadline should be a date in the future");
        crowdFunding.createCampaign(
            100 * 1e6,
            block.timestamp - 1, // Past deadline
            keccak256("past_deadline_test")
        );
    }

    function test_DonateToCampaign_FailsAfterDeadline() public {
        vm.prank(campaignOwner);
        uint256 campaignId = crowdFunding.createCampaign(
            100 * 1e6,
            block.timestamp + 1 days,
            keccak256("expired_campaign_test")
        );

        // Fast forward past deadline
        vm.warp(block.timestamp + 2 days);

        // Try to donate after deadline
        vm.startPrank(donor);
        usdc.approve(address(crowdFunding), 50 * 1e6);
        vm.expectRevert("Campaign has ended");
        crowdFunding.donateToCampaign(campaignId, 50 * 1e6);
        vm.stopPrank();
    }
}