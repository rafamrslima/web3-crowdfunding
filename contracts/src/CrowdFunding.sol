// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

contract CrowdFunding {
    struct Campaign {
        address owner;
        string title;
        string description;
        uint256 target;
        uint256 deadline;
        uint256 amountCollected;
        string image;
        address[] donators;
        uint256[] donations;
        bool withdrawn;
    }

    event CampaignCreated(
        uint256 indexed id,
        address indexed owner,
        string title,
        uint256 targetWei,
        uint256 deadline
    );

    event DonationReceived(
        uint256 indexed campaignId,
        address indexed receiver,
        address indexed donor,
        uint256 amountWei
    );

    event CampaignWithdrawn(
        uint256 indexed campaignId,
        address indexed owner,
        uint256 amountWei,
        uint256 timestamp
    );

    mapping(uint256 => Campaign) public campaigns;
    uint256 public numberOfCampaigns = 0;

    function createCampaign(address _owner, string memory _title, string memory _description,
    uint256 _target, uint256 _deadline, string memory _image) public returns (uint256) {
        uint256 id = numberOfCampaigns;
        Campaign storage campaign = campaigns[id];
        require(_deadline > block.timestamp, "The deadline should be a date in the future");

        campaign.owner = _owner;
        campaign.title = _title;
        campaign.description = _description;
        campaign.target = _target;
        campaign.deadline = _deadline;
        campaign.amountCollected = 0;
        campaign.image = _image;
        campaign.withdrawn = false;

        emit CampaignCreated(id, campaign.owner, campaign.title, campaign.target, campaign.deadline);

        numberOfCampaigns++;
        return id;
    }

    function donateToCampaign(uint256 _id) public payable {
        require(_id < numberOfCampaigns, "Campaign does not exist");
        Campaign storage campaign = campaigns[_id];
        require(block.timestamp < campaign.deadline, "Campaign has ended");

        uint256 amount = msg.value;
        campaign.amountCollected = campaign.amountCollected + amount;
        campaign.donations.push(amount);
        campaign.donators.push(msg.sender);
        campaign.amountCollected = campaign.amountCollected + amount;

        emit DonationReceived(_id, campaign.owner, msg.sender, amount);
    }

    function withdraw(uint256 _idCampaign) external {
        require(_idCampaign < numberOfCampaigns, "Campaign does not exist");
        Campaign storage campaign = campaigns[_idCampaign];
        require(block.timestamp >= campaign.deadline, "Campaign is still ongoing");
        require(campaign.amountCollected >= campaign.target, "Campaign didn't reach the target");
        require(msg.sender == campaign.owner, "Withdraw should be done by the campaign owner");
        require(!campaign.withdrawn, "Withdraw already done.");

        campaign.withdrawn = true;
        (bool sent,) = payable(campaign.owner).call{value: campaign.amountCollected}("");
        require(sent, "Transfer failed.");
        
        emit CampaignWithdrawn(_idCampaign, campaign.owner, campaign.amountCollected, block.timestamp);
    }

    function refundDonators(uint256 _idCampaign) public payable {
        require(_idCampaign < numberOfCampaigns, "Campaign does not exist");
        //todo
    }

    function getDonators(uint256 _id) public view returns (address[] memory, uint256[] memory) {
        Campaign storage campaign = campaigns[_id];
        return (campaign.donators, campaign.donations);
    }

    function getCampaigns() public view returns (Campaign[] memory) {
        Campaign[] memory allCampaigns = new Campaign[](numberOfCampaigns);

        for (uint256 i = 0; i < numberOfCampaigns; i++) {
            Campaign storage item = campaigns[i];
            allCampaigns[i] = item;
        }

        return allCampaigns;
    }
}
