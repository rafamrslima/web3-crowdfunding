import { Link, useLocation } from 'react-router-dom';
import './App.css';

interface SidebarProps {
  account: `0x${string}` | null;
  onConnectWallet: () => void;
  onDisconnect: () => void;
}

export default function Sidebar({ account, onConnectWallet, onDisconnect }: SidebarProps) {
  const location = useLocation();

  const isActive = (path: string) => {
    return location.pathname === path;
  };

  return (
    <div className="sidebar">
      <div className="sidebar-header">
        <h2 className="sidebar-title">🚀 CrowdFund</h2>
      </div>
      
      <nav className="sidebar-nav">
        {!account ? (
          <>
            <button 
              onClick={onConnectWallet}
              className="sidebar-button sidebar-button-primary"
            >
              🦊 Connect Wallet
            </button>
            
            <Link 
              to="/" 
              className={`sidebar-link ${isActive('/') ? 'active' : ''}`}
            >
              🏠 Dashboard
            </Link>
            
            <Link 
              to="/create" 
              className={`sidebar-link ${isActive('/create') ? 'active' : ''}`}
            >
              ➕ Create Campaign
            </Link>
          </>
        ) : (
          <>
            <div className="sidebar-account">
              <div className="account-badge">
                ✅ {account.slice(0, 6)}...{account.slice(-4)}
              </div>
            </div>
            
            <Link 
              to="/" 
              className={`sidebar-link ${isActive('/') ? 'active' : ''}`}
            >
              🏠 Dashboard
            </Link>
            
            <Link 
              to="/create" 
              className={`sidebar-link ${isActive('/create') ? 'active' : ''}`}
            >
              ➕ Create Campaign
            </Link>
            
            <Link 
              to="/my-campaigns" 
              className={`sidebar-link ${isActive('/my-campaigns') ? 'active' : ''}`}
            >
              📋 My Campaigns
            </Link>
            
            <Link 
              to="/my-donations" 
              className={`sidebar-link ${isActive('/my-donations') ? 'active' : ''}`}
            >
              💝 My Donations
            </Link>
            
            <Link 
              to="/refunds" 
              className={`sidebar-link ${isActive('/refunds') ? 'active' : ''}`}
            >
              🔄 Refunds
            </Link>
            
            <button 
              onClick={onDisconnect}
              className="sidebar-button sidebar-button-danger"
            >
              🔌 Logout
            </button>
          </>
        )}
      </nav>
    </div>
  );
}