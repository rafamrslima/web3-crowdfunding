import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { formatEther, createWalletClient, custom, parseEther } from 'viem';
import type { Hex } from 'viem';
import { API_BASE, CHAIN_ID } from './config';
import type { Campaign } from './types';
import './App.css';

declare global { interface Window { ethereum?: any } }

export default function CampaignsPage() {
  const [campaigns, setCampaigns] = useState<Campaign[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // Donation states
  const [account, setAccount] = useState<`0x${string}` | null>(null);
  const [expandedDonation, setExpandedDonation] = useState<number | null>(null);
  const [donationAmounts, setDonationAmounts] = useState<Record<number, string>>({});
  const [donationLoading, setDonationLoading] = useState<Record<number, boolean>>({});
  const [donationSuccess, setDonationSuccess] = useState<Record<number, string | null>>({});
  const [donationError, setDonationError] = useState<Record<number, string | null>>({});

  // MetaMask wallet client
  const walletClient = window.ethereum
    ? createWalletClient({ chain: { id: CHAIN_ID } as any, transport: custom(window.ethereum) })
    : null;

  useEffect(() => {
    fetchCampaigns();
  }, []);

  const fetchCampaigns = async () => {
    try {
      setLoading(true);
      setError(null);
      
      const response = await fetch(`${API_BASE}/api/v1/campaigns`);
      
      if (!response.ok) {
        throw new Error(`Failed to fetch campaigns: ${response.status} ${response.statusText}`);
      }
      
      const data: Campaign[] = await response.json();
      setCampaigns(data);
      
    } catch (err) {
      console.error('Error fetching campaigns:', err);
      setError(err instanceof Error ? err.message : 'Failed to fetch campaigns');
    } finally {
      setLoading(false);
    }
  };

  const formatWeiToEth = (weiAmount: number): string => {
    try {
      // Convert to string to handle large numbers, then format
      return formatEther(BigInt(weiAmount.toString()));
    } catch {
      return '0';
    }
  };

  const formatDeadline = (unixTimestamp: number): string => {
    return new Date(unixTimestamp * 1000).toLocaleDateString();
  };

  const calculateProgress = (target: number, collected: number): number => {
    if (target === 0) return 0;
    return Math.min((collected / target) * 100, 100);
  };

  // Connect to MetaMask
  const connectWallet = async () => {
    if (!walletClient) {
      alert("MetaMask not found. Please install MetaMask extension.");
      return;
    }

    try {
      await window.ethereum.request({ method: "eth_requestAccounts" });

      const current = await walletClient.getChainId();
      if (current !== CHAIN_ID) {
        try {
          await window.ethereum.request({
            method: "wallet_switchEthereumChain",
            params: [{ chainId: "0x" + CHAIN_ID.toString(16) }],
          });
        } catch {
          await window.ethereum.request({
            method: "wallet_addEthereumChain",
            params: [{
              chainId: "0x" + CHAIN_ID.toString(16),
              chainName: "Anvil Local",
              nativeCurrency: { name: "ETH", symbol: "ETH", decimals: 18 },
              rpcUrls: ["http://127.0.0.1:8545"],
            }],
          });
        }
      }

      const [addr] = await walletClient.getAddresses();
      setAccount(addr);
    } catch (err) {
      console.error("Failed to connect wallet:", err);
      alert("Failed to connect wallet. Please try again.");
    }
  };

  // Toggle donation form for a specific campaign
  const toggleDonationForm = async (campaignIndex: number) => {
    if (expandedDonation === campaignIndex) {
      setExpandedDonation(null);
      return;
    }

    if (!account) {
      await connectWallet();
      if (!account) return; // If connection failed
    }

    setExpandedDonation(campaignIndex);
    // Clear any previous states for this campaign
    setDonationError(prev => ({ ...prev, [campaignIndex]: null }));
    setDonationSuccess(prev => ({ ...prev, [campaignIndex]: null }));
  };

  // Update donation amount for a specific campaign
  const updateDonationAmount = (campaignIndex: number, amount: string) => {
    setDonationAmounts(prev => ({ ...prev, [campaignIndex]: amount }));
  };

  // Send donation transaction
  const sendDonation = async (campaignIndex: number) => {
    if (!walletClient || !account) {
      alert("Please connect your wallet first");
      return;
    }

    const donationAmount = donationAmounts[campaignIndex];
    if (!donationAmount || parseFloat(donationAmount) <= 0) {
      setDonationError(prev => ({ ...prev, [campaignIndex]: "Please enter a valid donation amount" }));
      return;
    }

    setDonationLoading(prev => ({ ...prev, [campaignIndex]: true }));
    setDonationError(prev => ({ ...prev, [campaignIndex]: null }));
    setDonationSuccess(prev => ({ ...prev, [campaignIndex]: null }));

    try {
      // Convert ETH to wei for the API call
      const weiAmount = parseEther(donationAmount);
      
      const payload = {
        campaignId: campaignIndex,
        value: weiAmount.toString()
      };

      console.log("Creating unsigned donation transaction...", payload);

      // Call the unsigned donation endpoint
      const response = await fetch(`${API_BASE}/api/v1/donations/unsigned`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload)
      });

      if (!response.ok) {
        const errorData = await response.text();
        throw new Error(`API error: ${response.status} ${response.statusText} - ${errorData}`);
      }

      const unsignedTx = await response.json();
      console.log("Received unsigned transaction:", unsignedTx);

      // Send transaction through MetaMask
      const txHash = await walletClient.sendTransaction({
        account,
        to: unsignedTx.to as `0x${string}`,
        data: unsignedTx.data as Hex,
        value: unsignedTx.value ? BigInt(unsignedTx.value) : 0n,
        gas: unsignedTx.gas ? BigInt(unsignedTx.gas) : undefined,
        chain: undefined
      });

      console.log("✅ Donation transaction successful:", txHash);
      
      // Show success message
      setDonationSuccess(prev => ({ ...prev, [campaignIndex]: txHash }));
      
      // Clear the donation form
      setDonationAmounts(prev => ({ ...prev, [campaignIndex]: "" }));
      setExpandedDonation(null);
      
      // Refresh campaigns data to show updated amounts
      setTimeout(() => {
        fetchCampaigns();
      }, 2000); // Wait 2 seconds for blockchain confirmation

    } catch (err) {
      console.error("Donation failed:", err);
      const error = err as Error;
      
      let errorMessage = "Failed to send donation";
      if (error.message?.includes("User rejected") || error.message?.includes("rejected")) {
        errorMessage = "Transaction was rejected by user";
      } else if (error.message?.includes("insufficient funds")) {
        errorMessage = "Insufficient ETH balance for donation";
      } else if (error.message) {
        errorMessage = error.message;
      }
      
      setDonationError(prev => ({ ...prev, [campaignIndex]: errorMessage }));
    } finally {
      setDonationLoading(prev => ({ ...prev, [campaignIndex]: false }));
    }
  };

  if (loading) {
    return (
      <div className="app-container">
        <div style={{ textAlign: 'center', padding: '3rem' }}>
          <div className="loading-spinner" style={{ margin: '0 auto' }}></div>
          <p>Loading campaigns...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="app-container">
        <div className="message-box message-error">
          <h4 className="message-title">❌ Error Loading Campaigns</h4>
          <p className="message-text">{error}</p>
          <button onClick={fetchCampaigns} className="btn btn-primary" style={{ marginTop: '1rem' }}>
            Try Again
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="app-container">
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '2rem' }}>
        <h1 className="app-title" style={{ margin: 0 }}>🌟 Active Campaigns</h1>
        <Link to="/create" className="btn btn-success">
          ➕ Create Campaign
        </Link>
      </div>

      {campaigns.length === 0 ? (
        <div className="message-box message-error">
          <h4 className="message-title">No Campaigns Found</h4>
          <p className="message-text">
            There are no active campaigns at the moment. Be the first to create one!
          </p>
          <Link to="/create" className="btn btn-primary" style={{ marginTop: '1rem' }}>
            Create First Campaign
          </Link>
        </div>
      ) : (
        <div className="campaigns-grid">
          {campaigns.map((campaign, index) => (
            <div key={index} className="campaign-card">
              <div className="campaign-header">
                {campaign.Image && (
                  <img 
                    src={campaign.Image} 
                    alt={campaign.Title}
                    className="campaign-image"
                    onError={(e) => {
                      // Hide image if it fails to load
                      e.currentTarget.style.display = 'none';
                    }}
                  />
                )}
                <h3 className="campaign-title">{campaign.Title}</h3>
              </div>

              <div className="campaign-content">
                <p className="campaign-description">{campaign.Description}</p>
                
                <div className="campaign-stats">
                  <div className="stat-item">
                    <span className="stat-label">Target:</span>
                    <span className="stat-value">{formatWeiToEth(campaign.Target)} ETH</span>
                  </div>
                  
                  <div className="stat-item">
                    <span className="stat-label">Collected:</span>
                    <span className="stat-value">{formatWeiToEth(campaign.AmountCollected)} ETH</span>
                  </div>
                  
                  <div className="stat-item">
                    <span className="stat-label">Deadline:</span>
                    <span className="stat-value">{formatDeadline(campaign.Deadline)}</span>
                  </div>
                </div>

                <div className="progress-section">
                  <div className="progress-bar">
                    <div 
                      className="progress-fill" 
                      style={{ width: `${calculateProgress(campaign.Target, campaign.AmountCollected)}%` }}
                    ></div>
                  </div>
                  <span className="progress-text">
                    {calculateProgress(campaign.Target, campaign.AmountCollected).toFixed(1)}% funded
                  </span>
                </div>

                <div className="campaign-owner">
                  <span className="stat-label">Owner:</span>
                  <span className="wallet-address">{campaign.Owner.slice(0, 6)}...{campaign.Owner.slice(-4)}</span>
                </div>
              </div>

              <div className="campaign-actions">
                <button 
                  className="btn btn-primary"
                  onClick={() => toggleDonationForm(index)}
                  disabled={donationLoading[index]}
                >
                  {donationLoading[index] && <span className="loading-spinner"></span>}
                  {expandedDonation === index ? "❌ Cancel" : "💝 Donate"}
                </button>
              </div>

              {/* Donation Form - Expandable */}
              {expandedDonation === index && (
                <div className="donation-form">
                  <h4>💰 Donate to this Campaign</h4>
                  <div className="form-group">
                    <label className="form-label">Donation Amount (ETH)</label>
                    <input
                      type="number"
                      min="0"
                      step="0.01"
                      placeholder="Enter ETH amount (e.g., 0.1)"
                      value={donationAmounts[index] || ""}
                      onChange={(e) => updateDonationAmount(index, e.target.value)}
                      className="form-input"
                    />
                  </div>
                  
                  <div className="donation-actions">
                    <button 
                      onClick={() => sendDonation(index)}
                      disabled={donationLoading[index] || !donationAmounts[index]}
                      className={`btn ${donationLoading[index] ? 'btn-secondary' : 'btn-success'}`}
                    >
                      {donationLoading[index] && <span className="loading-spinner"></span>}
                      {donationLoading[index] ? 'Processing...' : '🚀 Send Donation'}
                    </button>
                  </div>

                  {/* Success Message */}
                  {donationSuccess[index] && (
                    <div className="message-box message-success">
                      <h4 className="message-title">✅ Donation Sent Successfully!</h4>
                      <p className="message-text">Your donation has been sent to the blockchain.</p>
                      <p className="message-text">
                        <strong>Transaction ID:</strong>
                        <code className="transaction-hash">{donationSuccess[index]}</code>
                      </p>
                    </div>
                  )}

                  {/* Error Message */}
                  {donationError[index] && (
                    <div className="message-box message-error">
                      <h4 className="message-title">❌ Donation Failed</h4>
                      <p className="message-text">{donationError[index]}</p>
                    </div>
                  )}
                </div>
              )}
            </div>
          ))}
        </div>
      )}

      <div style={{ textAlign: 'center', marginTop: '3rem', padding: '2rem', backgroundColor: 'var(--light-gray)', borderRadius: 'var(--border-radius)' }}>
        <h3>Ready to make a difference? 🚀</h3>
        <p>Create your own campaign and let the community support your cause!</p>
        <Link to="/create" className="btn btn-success">
          Start Your Campaign
        </Link>
      </div>
    </div>
  );
}