// types.ts - Type definitions for the application

export interface Campaign {
  Owner: string;
  Title: string;
  Description: string;
  Target: number; // USDC amount as number (converted from string)
  Deadline: number; // Unix timestamp
  AmountCollected: number; // USDC amount as number (converted from string)
  Image: string;
  Donators: string[];
  Donations: number[];
  Withdrawn: boolean; 
}

export interface UnsignedTransaction {
  to: string; // Contract address
  data: string; // Transaction data (hex)
  value: string; // Value in USDC (hex)
  gas: string; // Gas limit (hex)
}