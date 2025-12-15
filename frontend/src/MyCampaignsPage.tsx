
import { useState, useEffect, useCallback } from 'react';
import { useWallet } from './WalletContext';
import { API_BASE } from './config';
import type { Hex } from 'viem';
import './App.css';

interface UserCampaign {
  owner: string;
  title: string;
  description: string;
  target: string; // USDC amount as string from API
  deadline: string; // Unix timestamp as string from API
  image: string;
  amountCollected: number | null; // USDC amount collected, can be null
}

export default function MyCampaignsPage() {
  const { account, walletClient } = useWallet();
  const [campaigns, setCampaigns] = useState<UserCampaign[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  
  // Withdrawal states
  const [withdrawLoading, setWithdrawLoading] = useState<Record<number, boolean>>({});
  const [withdrawSuccess, setWithdrawSuccess] = useState<Record<number, string | null>>({});
  const [withdrawError, setWithdrawError] = useState<Record<number, string | null>>({});

  const fetchUserCampaigns = useCallback(async () => {
    if (!account) return;
    
    try {
      setLoading(true);
      setError(null);
      
      const response = await fetch(`${API_BASE}/api/v1/campaigns/owner/${account}`);
      
      if (!response.ok) {
        throw new Error(`Failed to fetch campaigns: ${response.status} ${response.statusText}`);
      }
      
      const data: UserCampaign[] = await response.json();
      setCampaigns(data);
      
    } catch (err) {
      console.error('Error fetching user campaigns:', err);
      setError(err instanceof Error ? err.message : 'Failed to fetch campaigns');
    } finally {
      setLoading(false);
    }
  }, [account]);

  // Fetch user's campaigns when account is available
  useEffect(() => {
    if (account) {
      fetchUserCampaigns();
    } else {
      setCampaigns([]);
    }
  }, [account, fetchUserCampaigns]);

  const formatToUsdc = (usdcAmount: string): string => {
    try {
      // USDC format: last 6 digits are cents (micro-units)
      // Example: "103000000" = $103.00
      const amount = parseInt(usdcAmount);
      const dollars = amount / 1000000; // Divide by 1,000,000 to convert from micro-units
      const formatted = dollars.toFixed(2);
      return formatted === '0.00' ? '0' : formatted.replace(/\.?0+$/, '');
    } catch {
      return '0';
    }
  };

  const formatDeadline = (unixTimestamp: string): string => {
    try {
      return new Date(parseInt(unixTimestamp) * 1000).toLocaleDateString();
    } catch {
      return 'Invalid date';
    }
  };

  const isDeadlinePassed = (unixTimestamp: string): boolean => {
    try {
      const deadline = parseInt(unixTimestamp) * 1000;
      return Date.now() > deadline;
    } catch {
      return false;
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

  const isTargetReached = (target: string, collected: number | null): boolean => {
    if (!collected) return false;
    try {
      const targetAmount = parseInt(target);
      return collected >= targetAmount;
    } catch {
      return false;
    }
  };

  const canWithdraw = (campaign: UserCampaign): boolean => {
    return isDeadlinePassed(campaign.deadline) && isTargetReached(campaign.target, campaign.amountCollected);
  };

  const getWithdrawButtonText = (campaign: UserCampaign, campaignIndex: number): string => {
    if (withdrawLoading[campaignIndex]) return 'Processing...';
    if (!campaign.amountCollected || campaign.amountCollected === 0) return '💰 No Funds Yet';
    if (!isDeadlinePassed(campaign.deadline)) return '⏰ Deadline Not Reached';
    if (!isTargetReached(campaign.target, campaign.amountCollected)) return '🎯 Target Not Reached';
    return '💰 Withdraw';
  };

  const getWithdrawButtonTitle = (campaign: UserCampaign): string => {
    if (!campaign.amountCollected || campaign.amountCollected === 0) {
      return 'No funds collected yet';
    }
    if (!isDeadlinePassed(campaign.deadline)) {
      return 'Withdrawal only available after campaign deadline';
    }
    if (!isTargetReached(campaign.target, campaign.amountCollected)) {
      return 'Withdrawal only available when target amount is reached';
    }
    return 'Withdraw funds from this campaign';
  };

  const handleWithdraw = async (campaignIndex: number) => {
    if (!walletClient || !account) {
      alert('Please connect your wallet first');
      return;
    }

    try {
      setWithdrawLoading(prev => ({ ...prev, [campaignIndex]: true }));
      setWithdrawError(prev => ({ ...prev, [campaignIndex]: null }));
      setWithdrawSuccess(prev => ({ ...prev, [campaignIndex]: null }));

      console.log(`Creating withdraw transaction for campaign ${campaignIndex}...`);
      
      // Call API to get unsigned transaction
      const response = await fetch(`${API_BASE}/api/v1/campaigns/withdraw/${campaignIndex}`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' }
      });

      if (!response.ok) {
        const errorData = await response.text();
        throw new Error(`API error: ${response.status} ${response.statusText} - ${errorData}`);
      }

      const unsignedTx = await response.json();
      console.log('Received unsigned withdraw transaction:', unsignedTx);

      // Validate transaction data
      if (!unsignedTx.to || !unsignedTx.data || !unsignedTx.gas) {
        throw new Error('Invalid transaction data received from server');
      }

      console.log('Requesting MetaMask signature for withdrawal...');
      
      // Send transaction through MetaMask
      const txHash = await walletClient.sendTransaction({
        account,
        to: unsignedTx.to as `0x${string}`,
        data: unsignedTx.data as Hex,
        value: BigInt(unsignedTx.value || '0x0'),
        gas: BigInt(unsignedTx.gas),
        chain: undefined
      });

      console.log('✅ Withdrawal transaction successful:', txHash);
      setWithdrawSuccess(prev => ({ ...prev, [campaignIndex]: txHash }));
      
      // Refresh campaigns data after successful withdrawal
      setTimeout(() => {
        fetchUserCampaigns();
      }, 3000); // Wait 3 seconds for blockchain confirmation

    } catch (err) {
      console.error('Withdrawal failed:', err);
      const error = err as Error;
      
      let errorMessage = 'Failed to withdraw funds';
      if (error.message?.includes('User rejected') || error.message?.includes('rejected')) {
        errorMessage = 'Transaction was rejected by user';
      } else if (error.message?.includes('insufficient funds')) {
        errorMessage = 'Insufficient ETH balance for transaction fees';
      } else if (error.message) {
        errorMessage = error.message;
      }
      
      setWithdrawError(prev => ({ ...prev, [campaignIndex]: errorMessage }));
    } finally {
      setWithdrawLoading(prev => ({ ...prev, [campaignIndex]: false }));
    }
  };

  if (!account) {
    return (
      <div className="app-container">
        <div className="page-header">
          <h1 className="page-title">📋 My Campaigns</h1>
        </div>
        <div className="message-box message-error">
          <h4 className="message-title">🔒 Wallet Not Connected</h4>
          <p className="message-text">
            Please connect your wallet using the sidebar to view your campaigns.
          </p>
        </div>
      </div>
    );
  }

  if (loading) {
    return (
      <div className="app-container">
        <div className="page-header">
          <h1 className="page-title">📋 My Campaigns</h1>
          <div className="account-info">
            Connected: {account.slice(0, 6)}...{account.slice(-4)}
          </div>
        </div>
        <div style={{ textAlign: 'center', padding: '3rem' }}>
          <div className="loading-spinner" style={{ margin: '0 auto' }}></div>
          <p>Loading your campaigns...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="app-container">
        <div className="page-header">
          <h1 className="page-title">📋 My Campaigns</h1>
          <div className="account-info">
            Connected: {account.slice(0, 6)}...{account.slice(-4)}
          </div>
        </div>
        <div className="message-box message-error">
          <h4 className="message-title">❌ Error Loading Campaigns</h4>
          <p className="message-text">{error}</p>
          <button onClick={fetchUserCampaigns} className="btn btn-primary" style={{ marginTop: '1rem' }}>
            Try Again
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="app-container">
      <div className="page-header">
        <h1 className="page-title">📋 My Campaigns</h1>
        <div className="account-info">
          Connected: {account.slice(0, 6)}...{account.slice(-4)}
        </div>
      </div>

      {campaigns.length === 0 ? (
        <div className="message-box" style={{ backgroundColor: '#f8f9fa', border: '1px solid #dee2e6' }}>
          <h4 className="message-title">📝 No Campaigns Created</h4>
          <p className="message-text">
            You haven't created any campaigns yet. Start making a difference by creating your first campaign!
          </p>
          <a href="/create" className="btn btn-primary" style={{ marginTop: '1rem', textDecoration: 'none' }}>
            Create First Campaign
          </a>
        </div>
      ) : (
        <>
          <div style={{ marginBottom: '2rem', padding: '1rem', backgroundColor: 'var(--light-gray)', borderRadius: 'var(--border-radius)' }}>
            <p style={{ margin: '0 0 0.5rem 0', color: 'var(--text-muted)' }}>
              <strong>Total Campaigns:</strong> {campaigns.length}
            </p>
            <p style={{ margin: 0, fontSize: '0.9rem', color: 'var(--text-muted)' }}>
              💡 <strong>Withdrawal Note:</strong> You can only withdraw funds after the campaign deadline has passed AND the target amount has been reached.
            </p>
          </div>
          
          <div className="campaigns-grid">
            {campaigns.map((campaign, index) => (
              <div key={index} className="campaign-card">
                <div className="campaign-header">
                  {campaign.image && (
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
                    <div className="campaign-id">#{index}</div>
                    <h3 className="campaign-title">
                      {campaign.title || 'Untitled Campaign'}
                    </h3>
                  </div>
                </div>

                <div className="campaign-content">
                  <p className="campaign-description">
                    {campaign.description || 'No description provided'}
                  </p>
                  
                  <div className="campaign-stats">
                    <div className="stat-item">
                      <span className="stat-label">Target:</span>
                      <span className="stat-value">${formatToUsdc(campaign.target)} USDC</span>
                    </div>
                    
                    <div className="stat-item">
                      <span className="stat-label">Collected:</span>
                      <span className="stat-value">
                        ${campaign.amountCollected ? formatToUsdc(campaign.amountCollected.toString()) : '0'} USDC
                      </span>
                    </div>
                    
                    <div className="stat-item">
                      <span className="stat-label">Deadline:</span>
                      <span className={`stat-value ${isDeadlinePassed(campaign.deadline) ? 'text-danger' : ''}`}>
                        {formatDeadline(campaign.deadline)}
                        {isDeadlinePassed(campaign.deadline) && ' (Expired)'}
                      </span>
                    </div>
                    
                    <div className="stat-item">
                      <span className="stat-label">Status:</span>
                      <span className={`stat-value ${isDeadlinePassed(campaign.deadline) ? 'text-danger' : 'text-success'}`}>
                        {isDeadlinePassed(campaign.deadline) ? '⏰ Ended' : '✅ Active'}
                      </span>
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
                    className={`btn ${canWithdraw(campaign) ? 'btn-success' : 'btn-secondary'}`}
                    disabled={!canWithdraw(campaign) || withdrawLoading[index]}
                    onClick={() => handleWithdraw(index)}
                    title={getWithdrawButtonTitle(campaign)}
                  >
                    {withdrawLoading[index] && <span className="loading-spinner"></span>}
                    {getWithdrawButtonText(campaign, index)}
                  </button>
                </div>

                {/* Withdrawal Success Message */}
                {withdrawSuccess[index] && (
                  <div className="message-box message-success" style={{ marginTop: '1rem' }}>
                    <h4 className="message-title">✅ Withdrawal Successful!</h4>
                    <p className="message-text">Funds have been withdrawn from your campaign.</p>
                    <p className="message-text">
                      <strong>Transaction ID:</strong>
                      <code className="transaction-hash">{withdrawSuccess[index]}</code>
                    </p>
                  </div>
                )}

                {/* Withdrawal Error Message */}
                {withdrawError[index] && (
                  <div className="message-box message-error" style={{ marginTop: '1rem' }}>
                    <h4 className="message-title">❌ Withdrawal Failed</h4>
                    <p className="message-text">{withdrawError[index]}</p>
                  </div>
                )}
              </div>
            ))}
          </div>
        </>
      )}
    </div>
  );
}