// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract CrowdFunding {

    IERC20 public usdc;

    constructor(address _usdc) {
      require(_usdc != address(0), "invalid USDC");
      usdc = IERC20(_usdc);
    }

    struct Campaign {
        address owner;
        uint256 target;
        uint256 deadline;
        uint256 amountCollected;
        bool withdrawn;
    }

    event CampaignCreated(
        uint256 indexed id,
        address indexed owner,
        bytes32 indexed creationId,
        uint256 target,
        uint256 deadline
    );

    event DonationReceived(
        uint256 indexed campaignId,
        address indexed donor,
        uint256 amount
    );

    event FundsWithdrawn(
        uint256 indexed campaignId,
        address indexed owner,
        uint256 amount
    );

    event DonationRefunded(
        uint256 indexed campaignId,
        address indexed donor,
        uint256 totalContributed
    );

    mapping(uint256 => Campaign) public campaigns;
    mapping(uint256 => mapping(address => uint256)) public contributions; //campaignId -> (donator -> amount)
    uint256 public numberOfCampaigns = 0;

    function createCampaign(uint256 _target, uint256 _deadline, bytes32 _creationId) public returns (uint256) {
        uint256 id = numberOfCampaigns;
        Campaign storage campaign = campaigns[id];
        require(_deadline > block.timestamp, "The deadline should be a date in the future");
        require(_deadline >= block.timestamp + 7 days, "Campaign must run for at least 7 days");
        require(_target <= 10_000_000 * 1e6, "Target exceeds maximum allowed");

        campaign.owner = msg.sender;
        campaign.target = _target;
        campaign.deadline = _deadline;
        campaign.amountCollected = 0;
        campaign.withdrawn = false;

        emit CampaignCreated(id, campaign.owner, _creationId, campaign.target, campaign.deadline);

        numberOfCampaigns++;
        return id;
    }

    function donateToCampaign(uint256 _id, uint256 _amount) external {
        require(_id < numberOfCampaigns, "Campaign does not exist");
        Campaign storage campaign = campaigns[_id];
        require(block.timestamp < campaign.deadline, "Campaign has ended");
        require(_amount > 0, "Invalid amount");
        require(campaign.owner != msg.sender, "Owner can't donate");

        bool ok = usdc.transferFrom(msg.sender, address(this), _amount);
        require(ok, "USDC transfer failed");
        campaign.amountCollected += _amount;
        contributions[_id][msg.sender] += _amount;

        emit DonationReceived(_id, msg.sender, _amount);
    }

    function withdraw(uint256 _idCampaign) external {
        require(_idCampaign < numberOfCampaigns, "Campaign does not exist");
        Campaign storage campaign = campaigns[_idCampaign];
        require(block.timestamp > campaign.deadline, "Campaign is still ongoing");
        require(campaign.amountCollected >= campaign.target, "Campaign didn't reach the target");
        require(msg.sender == campaign.owner, "Withdraw should be done by the campaign owner");
        require(!campaign.withdrawn, "Withdraw already done.");

        campaign.withdrawn = true;
        bool ok = usdc.transfer(campaign.owner, campaign.amountCollected);
        require(ok, "Transfer failed.");
        
        emit FundsWithdrawn(_idCampaign, campaign.owner, campaign.amountCollected);
    }

    function refundDonor(uint256 _idCampaign) public {
        require(_idCampaign < numberOfCampaigns, "Campaign does not exist");
        Campaign storage campaign = campaigns[_idCampaign];
        require(campaign.deadline < block.timestamp, "Campaign is not ended yet");
        require(campaign.amountCollected < campaign.target, "Campaign goal was reached, no refund available");

        uint256 totalContributed = contributions[_idCampaign][msg.sender];
        require(totalContributed > 0, "No donation found");

        contributions[_idCampaign][msg.sender] = 0;

        bool ok = usdc.transfer(msg.sender, totalContributed);
        require(ok, "Refund failed");
        emit DonationRefunded(_idCampaign, msg.sender, totalContributed);
    }

    function getCampaignsTotal() public view returns (uint256 total) {
        return numberOfCampaigns;
    }
}
