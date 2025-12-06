import { createContext, useContext } from 'react';
import type { WalletClient } from 'viem';

interface WalletContextType {
  account: `0x${string}` | null;
  walletClient: WalletClient | null;
  connectWallet: () => Promise<void>;
  disconnectWallet: () => void;
}

export const WalletContext = createContext<WalletContextType | null>(null);

export function useWallet() {
  const context = useContext(WalletContext);
  if (!context) {
    throw new Error('useWallet must be used within a WalletProvider');
  }
  return context;
}