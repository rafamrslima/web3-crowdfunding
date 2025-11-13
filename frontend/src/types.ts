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
}