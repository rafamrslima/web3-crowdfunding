
import './App.css';

interface MyCampaignsPageProps {
  account: `0x${string}` | null;
}

export default function MyCampaignsPage({ account }: MyCampaignsPageProps) {
  return (
    <div className="page-container">
      <div className="page-header">
        <h1 className="page-title">📋 My Campaigns</h1>
        {account && (
          <div className="account-info">
            Connected: {account.slice(0, 6)}...{account.slice(-4)}
          </div>
        )}
      </div>
      
      <div className="blank-page">
        <div className="blank-content">
          <h3>🚧 Coming Soon</h3>
          <p>This page will display campaigns created by your wallet.</p>
          {!account && (
            <p className="text-muted">Please connect your wallet to view your campaigns.</p>
          )}
        </div>
      </div>
    </div>
  );
}