import { useState } from 'react';
import { createWalletClient, custom } from 'viem';
import { CHAIN_ID } from './config';
import Sidebar from './Sidebar';
import { WalletContext } from './WalletContext';
import './App.css';

declare global { interface Window { ethereum?: any } }

interface LayoutProps {
  children: React.ReactNode;
}

export default function Layout({ children }: LayoutProps) {
  const [account, setAccount] = useState<`0x${string}` | null>(null);

  // MetaMask wallet client
  const walletClient = window.ethereum
    ? createWalletClient({ 
        chain: { id: CHAIN_ID, name: 'Anvil', nativeCurrency: { decimals: 18, name: 'ETH', symbol: 'ETH' }, rpcUrls: { default: { http: ['http://127.0.0.1:8545'] } } }, 
        transport: custom(window.ethereum) 
      })
    : null;

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

  const disconnectWallet = () => {
    setAccount(null);
  };

  return (
    <WalletContext.Provider value={{
      account,
      walletClient,
      connectWallet,
      disconnectWallet
    }}>
      <div className="app-layout">
        <Sidebar 
          account={account}
          onConnectWallet={connectWallet}
          onDisconnect={disconnectWallet}
        />
        <div className="main-content">
          <div className="content-wrapper">
            {children}
          </div>
        </div>
      </div>
    </WalletContext.Provider>
  );
}