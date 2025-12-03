// types.ts - Type definitions for the application

export interface Campaign {
  Owner: string;
  Title: string;
  Description: string;
  Target: number; // Wei amount as number
  Deadline: number; // Unix timestamp
  AmountCollected: number; // Wei amount as number
  Image: string;
  Donators: string[];
  Donations: number[];
  Withdrawn: boolean; 
}

export interface UnsignedTransaction {
  to: string; // Contract address
  data: string; // Transaction data (hex)
  value: string; // Value in wei (hex)
  gas: string; // Gas limit (hex)
}