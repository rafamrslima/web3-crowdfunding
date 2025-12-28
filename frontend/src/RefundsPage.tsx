import { useState, useEffect, useCallback } from 'react';
import { createPublicClient, http } from 'viem';
import { anvil } from 'viem/chains';
import { API_BASE } from './config';
import { useWallet } from './WalletContext';
import { getRawContractError } from './contractErrors';
import './App.css';

interface RefundableDonation {
  donor: string;
  campaignId: string;
  title: string;
  description: string;
  createdAt: string;
  image: string;
  amount: number;
}

interface AvailableRefund {
  donation: RefundableDonation;
  target: string;
  deadline: string;
  amountCollected: string;
}

export default function RefundsPage() {
  const { account } = useWallet();
  const [refunds, setRefunds] = useState<AvailableRefund[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [processingRefunds, setProcessingRefunds] = useState<Set<string>>(new Set());

  // Format USDC amount for display
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
  const formatDate = (timestamp: string): string => {
    try {
      const date = new Date(parseInt(timestamp) * 1000);
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

  // Calculate progress percentage
  const calculateProgress = (target: string, collected: string): number => {
    try {
      const targetAmount = parseInt(target);
      const collectedAmount = parseInt(collected);
      if (targetAmount === 0) return 0;
      return Math.min((collectedAmount / targetAmount) * 100, 100);
    } catch {
      return 0;
    }
  };

  // Fetch available refunds
  const fetchRefunds = useCallback(async () => {
    if (!account) return;

    try {
      setLoading(true);
      setError(null);

      const response = await fetch(`${API_BASE}/api/v1/campaigns/refunds/${account}`);

      if (!response.ok) {
        if (response.status === 404) {
          setRefunds([]);
          return;
        }
        throw new Error(`Failed to fetch refunds: ${response.status} ${response.statusText}`);
      }

      const data: AvailableRefund[] = await response.json();
      setRefunds(data);

    } catch (err) {
      console.error('Error fetching refunds:', err);
      setError(err instanceof Error ? err.message : 'Failed to fetch available refunds');
    } finally {
      setLoading(false);
    }
  }, [account]);

  // Handle refund request
  const handleRefund = async (campaignId: string) => {
    if (!account || !window.ethereum) {
      setError('Please connect your wallet to request a refund');
      return;
    }

    try {
      setProcessingRefunds(prev => new Set(prev).add(campaignId));
      setError(null);

      // Call the refund API
      const response = await fetch(`${API_BASE}/api/v1/campaigns/refund/${campaignId}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(`Failed to prepare refund: ${errorText}`);
      }

      // The response is the transaction object directly
      const transaction = await response.json();

      // Send transaction using MetaMask
      const txHash = await window.ethereum.request({
        method: 'eth_sendTransaction',
        params: [{
          from: account,
          to: transaction.to,
          data: transaction.data,
          value: transaction.value,
          gas: transaction.gas,
        }],
      });

      // Wait for transaction confirmation
      const publicClient = createPublicClient({
        chain: anvil,
        transport: http('http://127.0.0.1:8545'),
      });

      const receipt = await publicClient.waitForTransactionReceipt({
        hash: txHash,
      });

      if (receipt.status === 'success') {
        alert('🎉 Refund successful! Your USDC has been returned to your wallet.');
        // Refresh refunds to update the list
        await fetchRefunds();
      } else {
        throw new Error('Refund transaction failed on blockchain');
      }

    } catch (error) {
      console.error('Refund error:', error);
      const errorMessage = getRawContractError(error) || 
        (error instanceof Error ? error.message : 'Failed to process refund');
      setError(`❌ Refund failed: ${errorMessage}`);
    } finally {
      setProcessingRefunds(prev => {
        const newSet = new Set(prev);
        newSet.delete(campaignId);
        return newSet;
      });
    }
  };

  useEffect(() => {
    if (account) {
      fetchRefunds();
    } else {
      setRefunds([]);
      setError(null);
    }
  }, [account, fetchRefunds]);

  if (!account) {
    return (
      <div className="app-container">
        <div className="page-header">
          <h1 className="page-title">🔄 Available Refunds</h1>
        </div>
        
        <div className="message-box message-error">
          <h4 className="message-title">🔒 Wallet Required</h4>
          <p className="message-text">
            Please connect your wallet to view your available refunds.
          </p>
        </div>
      </div>
    );
  }

  if (loading) {
    return (
      <div className="app-container">
        <div className="page-header">
          <h1 className="page-title">🔄 Available Refunds</h1>
          <div className="account-info">
            Connected: {account.slice(0, 6)}...{account.slice(-4)}
          </div>
        </div>
        
        <div style={{ textAlign: 'center', padding: '3rem' }}>
          <div className="loading-spinner" style={{ margin: '0 auto' }}></div>
          <p>Loading available refunds...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="app-container">
        <div className="page-header">
          <h1 className="page-title">🔄 Available Refunds</h1>
          <div className="account-info">
            Connected: {account.slice(0, 6)}...{account.slice(-4)}
          </div>
        </div>
        
        <div className="message-box message-error">
          <h4 className="message-title">❌ Error Loading Refunds</h4>
          <p className="message-text">{error}</p>
          <button onClick={fetchRefunds} className="btn btn-primary" style={{ marginTop: '1rem' }}>
            Try Again
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="app-container">
      <div className="page-header">
        <h1 className="page-title">🔄 Available Refunds</h1>
        <div className="account-info">
          Connected: {account.slice(0, 6)}...{account.slice(-4)}
        </div>
      </div>

      {refunds.length === 0 ? (
        <div className="message-box" style={{ backgroundColor: '#f8f9fa', border: '1px solid #dee2e6' }}>
          <h4 className="message-title">✅ No Refunds Available</h4>
          <p className="message-text">
            Great news! You don't have any failed campaigns eligible for refunds. 
            All campaigns you've supported either reached their targets or are still active.
          </p>
        </div>
      ) : (
        <>
          {/* Summary Section */}
          <div style={{ 
            marginBottom: '2rem', 
            padding: '1.5rem', 
            backgroundColor: '#fff3cd', 
            border: '1px solid #ffecb5',
            borderRadius: 'var(--border-radius)',
            textAlign: 'center'
          }}>
            <div style={{ fontSize: '1.2rem', fontWeight: 'bold', color: '#856404', marginBottom: '0.5rem' }}>
              ⚠️ {refunds.length} Campaign{refunds.length !== 1 ? 's' : ''} Available for Refund
            </div>
            <p style={{ margin: 0, color: '#856404', fontSize: '0.95rem' }}>
              These campaigns didn't reach their funding targets by the deadline. You can claim your donations back.
            </p>
          </div>

          {/* Refunds List */}
          <div className="refunds-list">
            {refunds.map((refund, index) => (
              <div key={index} className="refund-card" style={{
                backgroundColor: 'var(--white)',
                border: '1px solid var(--border-gray)',
                borderRadius: 'var(--border-radius)',
                padding: '1.5rem',
                marginBottom: '1.5rem',
                boxShadow: 'var(--box-shadow)',
                transition: 'var(--transition)'
              }}>
                <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', gap: '1rem' }}>
                  <div style={{ flex: 1 }}>
                    <div style={{ display: 'flex', alignItems: 'center', gap: '0.75rem', marginBottom: '0.75rem' }}>
                      <div className="campaign-id">#{refund.donation.campaignId}</div>
                      <h3 style={{ margin: 0, fontSize: '1.3rem', fontWeight: '600', color: 'var(--text-dark)' }}>
                        {refund.donation.title || 'Untitled Campaign'}
                      </h3>
                    </div>
                    
                    {refund.donation.description && (
                      <p style={{ 
                        margin: '0 0 1rem 0', 
                        color: 'var(--text-muted)', 
                        fontSize: '0.95rem',
                        lineHeight: '1.5'
                      }}>
                        {refund.donation.description}
                      </p>
                    )}

                    {/* Campaign Stats */}
                    <div style={{ 
                      display: 'flex', 
                      gap: '1.5rem',
                      fontSize: '0.9rem',
                      color: 'var(--text-muted)',
                      marginBottom: '1rem'
                    }}>
                      <span>🎯 Target: ${formatToUsdc(parseInt(refund.target))}</span>
                      <span>💰 Raised: ${formatToUsdc(parseInt(refund.amountCollected))}</span>
                      <span>⏰ Deadline: {formatDate(refund.deadline)}</span>
                    </div>

                    {/* Progress Bar */}
                    <div style={{ marginBottom: '1rem' }}>
                      <div style={{
                        width: '100%',
                        height: '8px',
                        backgroundColor: 'var(--light-gray)',
                        borderRadius: '4px',
                        overflow: 'hidden'
                      }}>
                        <div style={{
                          width: `${calculateProgress(refund.target, refund.amountCollected)}%`,
                          height: '100%',
                          backgroundColor: '#dc3545',
                          transition: 'width 0.3s ease'
                        }}></div>
                      </div>
                      <div style={{ 
                        fontSize: '0.8rem', 
                        color: 'var(--text-muted)', 
                        marginTop: '0.25rem' 
                      }}>
                        {calculateProgress(refund.target, refund.amountCollected).toFixed(1)}% funded (Failed)
                      </div>
                    </div>
                  </div>

                  {/* Refund Amount and Button */}
                  <div style={{ 
                    textAlign: 'right',
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'flex-end',
                    gap: '1rem'
                  }}>
                    <div>
                      <div style={{ 
                        fontSize: '1.4rem', 
                        fontWeight: 'bold', 
                        color: 'var(--warning-orange)',
                        display: 'flex',
                        alignItems: 'center',
                        gap: '0.25rem'
                      }}>
                        💰 ${formatToUsdc(refund.donation.amount)}
                      </div>
                      <div style={{ 
                        fontSize: '0.8rem', 
                        color: 'var(--text-muted)',
                        backgroundColor: 'var(--light-gray)',
                        padding: '0.25rem 0.5rem',
                        borderRadius: '12px',
                        marginTop: '0.25rem'
                      }}>
                        Your Donation
                      </div>
                    </div>

                    <button
                      onClick={() => handleRefund(refund.donation.campaignId)}
                      disabled={processingRefunds.has(refund.donation.campaignId)}
                      className="btn btn-warning"
                      style={{ 
                        minWidth: '140px',
                        fontSize: '0.9rem'
                      }}
                    >
                      {processingRefunds.has(refund.donation.campaignId) ? (
                        <>
                          <span className="loading-spinner" style={{ marginRight: '0.5rem' }}></span>
                          Processing...
                        </>
                      ) : (
                        '🔄 Claim Refund'
                      )}
                    </button>
                  </div>
                </div>

                {/* Campaign Image */}
                {refund.donation.image && refund.donation.image.trim() !== '' && (
                  <div style={{ marginTop: '1.5rem' }}>
                    <img 
                      src={refund.donation.image}
                      alt={refund.donation.title || 'Campaign Image'}
                      style={{
                        width: '100%',
                        height: '200px',
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
    </div>
  );
}