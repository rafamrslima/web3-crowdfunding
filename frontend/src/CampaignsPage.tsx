import { useState, useEffect, useCallback } from 'react';
import { Link } from 'react-router-dom';
import type { Hex } from 'viem';
import { API_BASE } from './config';
import type { Campaign, UnsignedTransaction } from './types';
import { approveUSDC, getUSDCBalance, needsApproval } from './utils/usdcApproval';
import { useWallet } from './WalletContext';
import { useCategories } from './hooks/useCategories';
import './App.css';

export default function CampaignsPage() {
  const { account, walletClient } = useWallet();
  const { categories, loading: categoriesLoading } = useCategories();
  const [campaigns, setCampaigns] = useState<Campaign[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [searchTerm, setSearchTerm] = useState<string>('');
  const [categoryFilter, setCategoryFilter] = useState<number | null>(null);

  // Donation states
  const [usdcBalance, setUsdcBalance] = useState<string>('0');
  const [expandedDonation, setExpandedDonation] = useState<number | null>(null);
  const [donationAmounts, setDonationAmounts] = useState<Record<number, string>>({});
  const [donationLoading, setDonationLoading] = useState<Record<number, boolean>>({});
  const [approvalLoading, setApprovalLoading] = useState<Record<number, boolean>>({});
  const [donationSuccess, setDonationSuccess] = useState<Record<number, string | null>>({});
  const [donationError, setDonationError] = useState<Record<number, string | null>>({});

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
      // Ensure campaigns are sorted by ID (array index represents campaign ID)
      setCampaigns(data);
      
    } catch (err) {
      console.error('Error fetching campaigns:', err);
      setError(err instanceof Error ? err.message : 'Failed to fetch campaigns');
    } finally {
      setLoading(false);
    }
  };

  const formatToUsdc = (usdcAmount: number): string => {
    try {
      // USDC format: last 6 digits are cents (micro-units)
      // Example: 103000000 = $103.00
      const dollars = usdcAmount / 1000000; // Divide by 1,000,000 to convert from micro-units
      const formatted = dollars.toFixed(2);
      return formatted === '0.00' ? '0' : formatted.replace(/\.?0+$/, '');
    } catch {
      return '0';
    }
  };

  const formatDeadline = (unixTimestamp: string): string => {
    try {
      const timestamp = parseInt(unixTimestamp);
      return new Date(timestamp * 1000).toLocaleDateString();
    } catch {
      return 'Invalid Date';
    }
  };

  const calculateProgress = (target: string, collected: number | null): number => {
    if (!collected || collected === 0) return 0;
    try {
      const targetAmount = parseInt(target);
      if (targetAmount === 0) return 0;
      return Math.min((collected / targetAmount) * 100, 100);
    } catch {
      return 0;
    }
  };

  // Load USDC balance function
  const loadUSDCBalance = useCallback(async (address: string) => {
    if (!walletClient) return;
    
    try {
      const balance = await getUSDCBalance(walletClient, address);
      setUsdcBalance(balance);
    } catch (err) {
      console.error("Failed to load USDC balance:", err);
      setUsdcBalance('0');
    }
  }, [walletClient]);

  // Load USDC balance when account changes
  useEffect(() => {
    if (account && walletClient) {
      loadUSDCBalance(account);
    } else {
      setUsdcBalance('0');
    }
  }, [account, walletClient, loadUSDCBalance]);

  // Toggle donation form for a specific campaign
  const toggleDonationForm = async (campaignIndex: number) => {
    if (expandedDonation === campaignIndex) {
      setExpandedDonation(null);
      return;
    }

    if (!account) {
      alert("Please connect your wallet first using the sidebar.");
      return;
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

  // Send donation transaction with USDC approval
  const sendDonation = async (campaignIndex: number) => {
    if (!walletClient || !account) {
      alert("Please connect your wallet first");
      return;
    }

    const donationAmount = donationAmounts[campaignIndex];
    if (!donationAmount || donationAmount.trim() === '' || parseFloat(donationAmount) <= 0 || isNaN(parseFloat(donationAmount))) {
      setDonationError(prev => ({ ...prev, [campaignIndex]: "Please enter a valid USDC amount (e.g., 10.5 or 50)" }));
      return;
    }

    // Check USDC balance
    const balance = parseFloat(usdcBalance);
    const amount = parseFloat(donationAmount);
    if (amount > balance) {
      setDonationError(prev => ({ ...prev, [campaignIndex]: `Insufficient USDC balance. You have $${balance} USDC` }));
      return;
    }

    setDonationLoading(prev => ({ ...prev, [campaignIndex]: true }));
    setDonationError(prev => ({ ...prev, [campaignIndex]: null }));
    setDonationSuccess(prev => ({ ...prev, [campaignIndex]: null }));

    try {
      // Step 1: Check if approval is needed and approve USDC spending
      const approvalNeeded = await needsApproval(walletClient, account, donationAmount);
      
      if (approvalNeeded) {
        console.log("Step 1: USDC approval needed, requesting approval...");
        setApprovalLoading(prev => ({ ...prev, [campaignIndex]: true }));
        
        const approvalTxHash = await approveUSDC(walletClient, account, donationAmount);
        console.log("✅ USDC approval transaction sent:", approvalTxHash);
        
        setApprovalLoading(prev => ({ ...prev, [campaignIndex]: false }));
        
        // Wait a moment for approval transaction to be processed
        await new Promise(resolve => setTimeout(resolve, 2000));
      } else {
        console.log("USDC approval not needed, proceeding with donation...");
      }

      // Step 2: Send donation transaction
      // Send USDC amount as string directly (e.g., "10.5" or "50")
      const payload = {
        campaignId: campaignIndex,
        value: donationAmount // Send USDC as string
      };

      console.log("Creating unsigned donation transaction...", payload);

      // Call the donation endpoint
      const response = await fetch(`${API_BASE}/api/v1/donations/create`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload)
      });

      if (!response.ok) {
        const errorData = await response.text();
        throw new Error(`API error: ${response.status} ${response.statusText} - ${errorData}`);
      }

      const unsignedTx: UnsignedTransaction = await response.json();
      console.log("Received unsigned transaction:", unsignedTx);

      // Validate and convert hex values
      if (!unsignedTx.to || !unsignedTx.data || !unsignedTx.value || !unsignedTx.gas) {
        throw new Error("Invalid transaction data received from server");
      }

      // Send transaction through MetaMask
      const txHash = await walletClient.sendTransaction({
        account,
        to: unsignedTx.to as `0x${string}`,
        data: unsignedTx.data as Hex,
        value: BigInt(unsignedTx.value), // Convert hex string to BigInt
        gas: BigInt(unsignedTx.gas), // Convert hex string to BigInt
        chain: undefined
      });

      console.log("✅ Donation transaction successful:", txHash);
      
      // Show success message
      setDonationSuccess(prev => ({ ...prev, [campaignIndex]: txHash }));
      
      // Clear the donation form
      setDonationAmounts(prev => ({ ...prev, [campaignIndex]: "" }));
      setExpandedDonation(null);
      
      // Refresh campaigns data and balance to show updated amounts
      setTimeout(() => {
        fetchCampaigns();
        if (account) loadUSDCBalance(account);
      }, 2000); // Wait 2 seconds for blockchain confirmation

    } catch (err) {
      console.error("Donation failed:", err);
      
      // Import and use the raw contract error parser
      const { getRawContractError } = await import('./contractErrors');
      
      const errorMessage = getRawContractError(err);
      
      setDonationError(prev => ({ ...prev, [campaignIndex]: errorMessage }));
    } finally {
      setDonationLoading(prev => ({ ...prev, [campaignIndex]: false }));
      setApprovalLoading(prev => ({ ...prev, [campaignIndex]: false }));
    }
  };

  // Filter campaigns based on search term and category
  const filteredCampaigns = campaigns.filter(campaign => {
    // Category filter
    if (categoryFilter !== null && campaign.categoryId !== categoryFilter) {
      return false;
    }
    
    // Search filter
    if (!searchTerm.trim()) return true;
    
    const searchLower = searchTerm.toLowerCase();
    const titleMatch = campaign.title?.toLowerCase().includes(searchLower);
    const descriptionMatch = campaign.description?.toLowerCase().includes(searchLower);
    
    return titleMatch || descriptionMatch;
  });

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
      <div className="page-header">
        <h1 className="page-title">🌟 Active Campaigns</h1>
        {campaigns !== null && campaigns.length > 0 && (
          <p className="page-subtitle">
            {filteredCampaigns.length} of {campaigns.length} campaign{campaigns.length !== 1 ? 's' : ''} 
            {(searchTerm.trim() || categoryFilter !== null) && filteredCampaigns.length !== campaigns.length ? ' match your filters' : ' available'}
          </p>
        )}
      </div>

      {/* Filter Controls */}
      {campaigns !== null && campaigns.length > 0 && (
        <div style={{ marginBottom: '2rem' }}>
          {/* Category Filter */}
          <div className="form-group" style={{ marginBottom: '1rem' }}>
            <label className="form-label">📂 Filter by Category</label>
            <div style={{ display: 'flex', gap: '0.5rem', alignItems: 'center' }}>
              {categoriesLoading ? (
                <div style={{ padding: '0.5rem', color: '#666' }}>Loading categories...</div>
              ) : (
                <>
                  <select
                    value={categoryFilter || ""}
                    onChange={(e) => {
                      const value = e.target.value;
                      setCategoryFilter(value === "" ? null : parseInt(value));
                    }}
                    className="form-input"
                    style={{ maxWidth: '600px' }}
                  >
                    <option value="">All Categories</option>
                    {categories.map((category) => (
                      <option key={category.id} value={category.id}>
                        {category.name}
                      </option>
                    ))}
                  </select>
                  {categoryFilter !== null && (
                    <button 
                      onClick={() => setCategoryFilter(null)}
                      className="btn btn-secondary"
                      style={{ padding: '0.5rem 1rem' }}
                    >
                      ✕ Clear
                    </button>
                  )}
                </>
              )}
            </div>
          </div>

          {/* Search Input */}
          <div className="form-group">
            <label className="form-label">🔍 Search Campaigns</label>
            <input
              type="text"
              placeholder="Search by title or description..."
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
              className="form-input"
              style={{ maxWidth: '600px' }}
            />
            {searchTerm.trim() && (
              <button 
                onClick={() => setSearchTerm('')}
                className="btn btn-secondary"
                style={{ marginLeft: '0.5rem', padding: '0.5rem 1rem' }}
              >
                ✕ Clear
              </button>
            )}
          </div>
        </div>
      )}

      {campaigns === null || campaigns.length === 0 ? (
        <div className="message-box message-error">
          <h4 className="message-title">No Campaigns Found</h4>
          <p className="message-text">
            There are no active campaigns at the moment. Be the first to create one!
          </p>
          <Link to="/create" className="btn btn-primary" style={{ marginTop: '1rem' }}>
            Create First Campaign
          </Link>
        </div>
      ) : filteredCampaigns.length === 0 ? (
        <div className="message-box" style={{ backgroundColor: '#fff3cd', border: '1px solid #ffc107' }}>
          <h4 className="message-title">🔍 No campaigns match your filters</h4>
          <p className="message-text">
            Try adjusting your search term or category filter to see more campaigns.
          </p>
          {(searchTerm.trim() || categoryFilter !== null) && (
            <button 
              onClick={() => {
                setSearchTerm('');
                setCategoryFilter(null);
              }}
              className="btn btn-primary"
              style={{ marginTop: '1rem' }}
            >
              Clear All Filters
            </button>
          )}
        </div>
      ) : (
        <div className="campaigns-grid">
          {filteredCampaigns.map((campaign) => {
            // Find the original index from the campaigns array
            const originalIndex = campaigns.indexOf(campaign);
            return (
            <div key={originalIndex} className="campaign-card">
              <div className="campaign-header">
                {campaign.image && campaign.image.trim() !== '' && (
                  <img 
                    src={campaign.image} 
                    alt={campaign.title || 'Campaign Image'}
                    className="campaign-image"
                    onError={(e) => {
                      // Hide image if it fails to load
                      e.currentTarget.style.display = 'none';
                    }}
                  />
                )}
                <div className="campaign-title-section">
                  <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', flexWrap: 'wrap', gap: '0.5rem' }}>
                    <div className="campaign-id">#{originalIndex}</div>
                    {campaign.categoryId && categories.find(c => c.id === campaign.categoryId) && (
                      <span style={{
                        fontSize: '0.75rem',
                        padding: '0.25rem 0.5rem',
                        backgroundColor: 'var(--primary-color)',
                        color: 'white',
                        borderRadius: '12px',
                        fontWeight: '500'
                      }}>
                        {categories.find(c => c.id === campaign.categoryId)?.name}
                      </span>
                    )}
                  </div>
                  <h3 className="campaign-title">
                    {campaign.title && campaign.title.trim() !== '' ? campaign.title : 'Untitled Campaign'}
                  </h3>
                </div>
              </div>

              <div className="campaign-content">
                <p className="campaign-description">
                  {campaign.description && campaign.description.trim() !== '' ? campaign.description : 'No description provided.'}
                </p>
                
                <div className="campaign-stats">
                  <div className="stat-item">
                    <span className="stat-label">Target:</span>
                    <span className="stat-value">${formatToUsdc(parseInt(campaign.target))} USDC</span>
                  </div>
                  
                  <div className="stat-item">
                    <span className="stat-label">Collected:</span>
                    <span className="stat-value">
                      ${campaign.amountCollected ? formatToUsdc(campaign.amountCollected) : '0'} USDC
                    </span>
                  </div>
                  
                  <div className="stat-item">
                    <span className="stat-label">Deadline:</span>
                    <span className="stat-value">{formatDeadline(campaign.deadline)}</span>
                  </div>
                </div>

                <div className="progress-section">
                  <div className="progress-bar">
                    <div 
                      className="progress-fill" 
                      style={{ width: `${calculateProgress(campaign.target, campaign.amountCollected)}%` }}
                    ></div>
                  </div>
                  <span className="progress-text">
                    {calculateProgress(campaign.target, campaign.amountCollected).toFixed(1)}% funded
                  </span>
                </div>

                <div className="campaign-owner">
                  <span className="stat-label">Owner:</span>
                  <span className="wallet-address">{campaign.owner.slice(0, 6)}...{campaign.owner.slice(-4)}</span>
                </div>
              </div>

              <div className="campaign-actions">
                <button 
                  className={`btn ${!account ? 'btn-secondary' : 'btn-primary'}`}
                  onClick={() => toggleDonationForm(originalIndex)}
                  disabled={!account || donationLoading[originalIndex] || approvalLoading[originalIndex]}
                  title={!account ? 'Please connect your wallet to donate' : ''}
                >
                  {(donationLoading[originalIndex] || approvalLoading[originalIndex]) && <span className="loading-spinner"></span>}
                  {!account ? "🔒 Connect Wallet to Donate" : (expandedDonation === originalIndex ? "❌ Cancel" : "💝 Donate")}
                </button>
              </div>

              {/* Donation Form - Expandable */}
              {expandedDonation === originalIndex && (
                <div className="donation-form">
                  <h4>💰 Donate to this Campaign</h4>
                  
                  {/* USDC Balance Display */}
                  {account && (
                    <div className="balance-info" style={{ 
                      backgroundColor: 'var(--light-gray)', 
                      padding: '1rem', 
                      borderRadius: 'var(--border-radius)', 
                      marginBottom: '1rem' 
                    }}>
                      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                        <span><strong>Your USDC Balance:</strong></span>
                        <span style={{ color: 'var(--success-color)', fontWeight: 'bold' }}>${usdcBalance} USDC</span>
                      </div>
                    </div>
                  )}
                  
                  <div className="form-group">
                    <label className="form-label">Donation Amount (USDC)</label>
                    <input
                      type="text"
                      placeholder="Enter USDC amount (e.g., 10.5 or 50)"
                      value={donationAmounts[originalIndex] || ""}
                      onChange={(e) => {
                        // Only allow numbers, dots, and basic validation
                        const value = e.target.value;
                        if (value === '' || /^\d*\.?\d*$/.test(value)) {
                          updateDonationAmount(originalIndex, value);
                        }
                      }}
                      className="form-input"
                    />
                  </div>
                  
                  <div className="donation-actions">
                    <button 
                      onClick={() => sendDonation(originalIndex)}
                      disabled={donationLoading[originalIndex] || approvalLoading[originalIndex] || !donationAmounts[originalIndex]}
                      className={`btn ${donationLoading[originalIndex] || approvalLoading[originalIndex] ? 'btn-secondary' : 'btn-success'}`}
                    >
                      {(donationLoading[originalIndex] || approvalLoading[originalIndex]) && <span className="loading-spinner"></span>}
                      {approvalLoading[originalIndex] ? '📝 Approving USDC...' : donationLoading[originalIndex] ? '💸 Processing Donation...' : '🚀 Send Donation'}
                    </button>
                  </div>

                  {/* Approval Info */}
                  {approvalLoading[originalIndex] && (
                    <div className="message-box" style={{ backgroundColor: '#e3f2fd', border: '1px solid #2196f3', marginTop: '1rem' }}>
                      <h4 style={{ color: '#1976d2', margin: '0 0 0.5rem 0' }}>🔐 USDC Approval Required</h4>
                      <p style={{ margin: 0, fontSize: '0.9rem', color: '#666' }}>
                        Please approve the USDC spending in MetaMask. This allows the crowdfunding contract to use your USDC tokens.
                      </p>
                    </div>
                  )}

                  {/* Success Message */}
                  {donationSuccess[originalIndex] && (
                    <div className="message-box message-success">
                      <h4 className="message-title">✅ Donation Sent Successfully!</h4>
                      <p className="message-text">Your donation has been sent to the blockchain.</p>
                      <p className="message-text">
                        <strong>Transaction ID:</strong>
                        <code className="transaction-hash">{donationSuccess[originalIndex]}</code>
                      </p>
                    </div>
                  )}

                  {/* Error Message */}
                  {donationError[originalIndex] && (
                    <div className="message-box message-error">
                      <h4 className="message-title">❌ Donation Failed</h4>
                      <p className="message-text">{donationError[originalIndex]}</p>
                    </div>
                  )}
                </div>
              )}
            </div>
            );
          })}
        </div>
      )}

      <div style={{ textAlign: 'center', marginTop: '3rem', padding: '2rem', backgroundColor: 'var(--light-gray)', borderRadius: 'var(--border-radius)' }}>
        <h3>Ready to make a difference? 🚀</h3>
        <p>Create your own campaign and let the community support your cause!</p>
        <p className="text-muted">Use the sidebar to create a new campaign!</p>
      </div>
    </div>
  );
}