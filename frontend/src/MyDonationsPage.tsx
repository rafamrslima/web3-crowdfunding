
import { useState, useEffect, useCallback } from 'react';
import { API_BASE } from './config';
import { useWallet } from './WalletContext';
import './App.css';

interface Donation {
  donor: string;
  campaignId: string;
  title: string;
  description: string;
  createdAt: string;
  image: string;
  amount: number;
}

export default function MyDonationsPage() {
  const { account } = useWallet();
  const [donations, setDonations] = useState<Donation[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  // Format USDC amount from micro-units to dollars
  const formatToUsdc = (usdcAmount: number): string => {
    try {
      const dollars = usdcAmount / 1000000; // Convert from micro-units
      const formatted = dollars.toFixed(2);
      return formatted === '0.00' ? '0' : formatted.replace(/\.?0+$/, '');
    } catch {
      return '0';
    }
  };

  // Format date for display
  const formatDate = (dateString: string): string => {
    try {
      const date = new Date(dateString);
      return date.toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
      });
    } catch {
      return 'Invalid Date';
    }
  };

  // Calculate total donations
  const getTotalDonated = (): string => {
    const total = donations.reduce((sum, donation) => sum + donation.amount, 0);
    return formatToUsdc(total);
  };

  // Get unique campaigns count
  const getUniqueCampaignsCount = (): number => {
    const uniqueCampaigns = new Set(donations.map(d => d.campaignId));
    return uniqueCampaigns.size;
  };

  // Fetch user donations
  const fetchDonations = useCallback(async () => {
    if (!account) return;

    try {
      setLoading(true);
      setError(null);

      const response = await fetch(`${API_BASE}/api/v1/donations/${account}`);

      if (!response.ok) {
        throw new Error(`Failed to fetch donations: ${response.status} ${response.statusText}`);
      }

      const data: Donation[] = await response.json();
      // Sort donations by date (newest first)
      const sortedData = data.sort((a, b) => new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime());
      setDonations(sortedData);

    } catch (err) {
      console.error('Error fetching donations:', err);
      setError(err instanceof Error ? err.message : 'Failed to fetch donations');
    } finally {
      setLoading(false);
    }
  }, [account]);

  useEffect(() => {
    if (account) {
      fetchDonations();
    } else {
      setDonations([]);
      setError(null);
    }
  }, [account, fetchDonations]);

  if (!account) {
    return (
      <div className="app-container">
        <div className="page-header">
          <h1 className="page-title">💝 My Donations</h1>
        </div>
        
        <div className="message-box message-error">
          <h4 className="message-title">🔒 Wallet Required</h4>
          <p className="message-text">
            Please connect your wallet to view your donation history.
          </p>
        </div>
      </div>
    );
  }

  if (loading) {
    return (
      <div className="app-container">
        <div className="page-header">
          <h1 className="page-title">💝 My Donations</h1>
          <div className="account-info">
            Connected: {account.slice(0, 6)}...{account.slice(-4)}
          </div>
        </div>
        
        <div style={{ textAlign: 'center', padding: '3rem' }}>
          <div className="loading-spinner" style={{ margin: '0 auto' }}></div>
          <p>Loading your donations...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="app-container">
        <div className="page-header">
          <h1 className="page-title">💝 My Donations</h1>
          <div className="account-info">
            Connected: {account.slice(0, 6)}...{account.slice(-4)}
          </div>
        </div>
        
        <div className="message-box message-error">
          <h4 className="message-title">❌ Error Loading Donations</h4>
          <p className="message-text">{error}</p>
          <button onClick={fetchDonations} className="btn btn-primary" style={{ marginTop: '1rem' }}>
            Try Again
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="app-container">
      <div className="page-header">
        <h1 className="page-title">💝 My Donations</h1>
        <div className="account-info">
          Connected: {account.slice(0, 6)}...{account.slice(-4)}
        </div>
      </div>

      {donations.length === 0 ? (
        <div className="message-box" style={{ backgroundColor: '#f8f9fa', border: '1px solid #dee2e6' }}>
          <h4 className="message-title">📝 No Donations Yet</h4>
          <p className="message-text">
            You haven't made any donations yet. Start supporting campaigns to see your donation history here!
          </p>
        </div>
      ) : (
        <>
          {/* Summary Section */}
          <div style={{ 
            marginBottom: '2rem', 
            padding: '1.5rem', 
            backgroundColor: 'var(--light-gray)', 
            borderRadius: 'var(--border-radius)',
            display: 'grid',
            gridTemplateColumns: 'repeat(auto-fit, minmax(200px, 1fr))',
            gap: '1rem'
          }}>
            <div style={{ textAlign: 'center' }}>
              <div style={{ fontSize: '2rem', fontWeight: 'bold', color: 'var(--success-green)' }}>
                ${getTotalDonated()}
              </div>
              <div style={{ color: 'var(--text-muted)', fontSize: '0.9rem' }}>Total Donated (USDC)</div>
            </div>
            <div style={{ textAlign: 'center' }}>
              <div style={{ fontSize: '2rem', fontWeight: 'bold', color: 'var(--primary-blue)' }}>
                {donations.length}
              </div>
              <div style={{ color: 'var(--text-muted)', fontSize: '0.9rem' }}>Total Donations</div>
            </div>
            <div style={{ textAlign: 'center' }}>
              <div style={{ fontSize: '2rem', fontWeight: 'bold', color: 'var(--warning-orange)' }}>
                {getUniqueCampaignsCount()}
              </div>
              <div style={{ color: 'var(--text-muted)', fontSize: '0.9rem' }}>Campaigns Supported</div>
            </div>
          </div>

          {/* Donations List */}
          <div className="donations-list">
            {donations.map((donation, index) => (
              <div key={index} className="donation-card" style={{
                backgroundColor: 'var(--white)',
                border: '1px solid var(--border-gray)',
                borderRadius: 'var(--border-radius)',
                padding: '1.5rem',
                marginBottom: '1rem',
                boxShadow: 'var(--box-shadow)',
                transition: 'var(--transition)'
              }}>
                <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', gap: '1rem' }}>
                  <div style={{ flex: 1 }}>
                    <div style={{ display: 'flex', alignItems: 'center', gap: '0.75rem', marginBottom: '0.75rem' }}>
                      <div className="campaign-id">#{donation.campaignId}</div>
                      <h3 style={{ margin: 0, fontSize: '1.25rem', fontWeight: '600', color: 'var(--text-dark)' }}>
                        {donation.title || 'Untitled Campaign'}
                      </h3>
                    </div>
                    
                    {donation.description && (
                      <p style={{ 
                        margin: '0 0 1rem 0', 
                        color: 'var(--text-muted)', 
                        fontSize: '0.95rem',
                        lineHeight: '1.5'
                      }}>
                        {donation.description}
                      </p>
                    )}

                    <div style={{ 
                      display: 'flex', 
                      alignItems: 'center', 
                      gap: '1rem',
                      fontSize: '0.9rem',
                      color: 'var(--text-muted)'
                    }}>
                      <span>🕒 {formatDate(donation.createdAt)}</span>
                    </div>
                  </div>

                  <div style={{ 
                    textAlign: 'right',
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'flex-end',
                    gap: '0.5rem'
                  }}>
                    <div style={{ 
                      fontSize: '1.5rem', 
                      fontWeight: 'bold', 
                      color: 'var(--success-green)',
                      display: 'flex',
                      alignItems: 'center',
                      gap: '0.25rem'
                    }}>
                      💰 ${formatToUsdc(donation.amount)}
                    </div>
                    <div style={{ 
                      fontSize: '0.8rem', 
                      color: 'var(--text-muted)',
                      backgroundColor: 'var(--light-gray)',
                      padding: '0.25rem 0.5rem',
                      borderRadius: '12px'
                    }}>
                      USDC
                    </div>
                  </div>
                </div>

                {donation.image && donation.image.trim() !== '' && (
                  <div style={{ marginTop: '1rem' }}>
                    <img 
                      src={donation.image}
                      alt={donation.title || 'Campaign Image'}
                      style={{
                        width: '100%',
                        height: '150px',
                        objectFit: 'cover',
                        borderRadius: 'var(--border-radius)',
                        border: '1px solid var(--border-gray)'
                      }}
                      onError={(e) => {
                        e.currentTarget.style.display = 'none';
                      }}
                    />
                  </div>
                )}
              </div>
            ))}
          </div>
        </>
      )}

      <div style={{ 
        textAlign: 'center', 
        marginTop: '3rem', 
        padding: '2rem', 
        backgroundColor: 'var(--light-gray)', 
        borderRadius: 'var(--border-radius)' 
      }}>
        <h3>💝 Thank you for your generosity! 🙏</h3>
        <p>Your donations help make dreams come true and create positive impact.</p>
        <p className="text-muted">Keep supporting amazing causes!</p>
      </div>
    </div>
  );
}