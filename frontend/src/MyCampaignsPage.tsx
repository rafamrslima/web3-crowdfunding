
import { useState, useEffect, useCallback } from 'react';
import { useWallet } from './WalletContext';
import { API_BASE } from './config';
import './App.css';

interface UserCampaign {
  owner: string;
  title: string;
  description: string;
  target: string; // USDC amount as string from API
  deadline: string; // Unix timestamp as string from API
  image: string;
  AmountCollected: number | null; // USDC amount collected, can be null
}

export default function MyCampaignsPage() {
  const { account } = useWallet();
  const [campaigns, setCampaigns] = useState<UserCampaign[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

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
            <p style={{ margin: 0, color: 'var(--text-muted)' }}>
              <strong>Total Campaigns:</strong> {campaigns.length}
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
                  <h3 className="campaign-title">
                    {campaign.title || `Campaign #${index + 1}`}
                  </h3>
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
                        ${campaign.AmountCollected ? formatToUsdc(campaign.AmountCollected.toString()) : '0'} USDC
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
                        style={{ width: `${calculateProgress(campaign.target, campaign.AmountCollected)}%` }}
                      ></div>
                    </div>
                    <span className="progress-text">
                      {calculateProgress(campaign.target, campaign.AmountCollected).toFixed(1)}% funded
                    </span>
                  </div>

                  <div className="campaign-owner">
                    <span className="stat-label">Owner:</span>
                    <span className="wallet-address">{campaign.owner.slice(0, 6)}...{campaign.owner.slice(-4)}</span>
                  </div>
                </div>

                <div className="campaign-actions">
                  <button 
                    className="btn btn-success"
                    disabled={isDeadlinePassed(campaign.deadline) || !campaign.AmountCollected || campaign.AmountCollected === 0}
                    title={
                      !campaign.AmountCollected || campaign.AmountCollected === 0 
                        ? 'No funds collected yet' 
                        : isDeadlinePassed(campaign.deadline) 
                          ? 'Campaign has ended' 
                          : 'Withdraw funds from this campaign'
                    }
                  >
                    {!campaign.AmountCollected || campaign.AmountCollected === 0 
                      ? '💰 No Funds Yet' 
                      : isDeadlinePassed(campaign.deadline) 
                        ? '⏰ Campaign Ended' 
                        : '💰 Withdraw'}
                  </button>
                </div>
              </div>
            ))}
          </div>
        </>
      )}
    </div>
  );
}