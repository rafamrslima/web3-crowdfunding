// types.ts - Type definitions for the application

export interface Category {
  id: number;
  name: string;
  slug: string;
  description: string;
}

export interface Campaign {
  owner: string;
  title: string;
  description: string;
  target: string; // USDC amount as string
  deadline: string; // Unix timestamp as string
  amountCollected: number | null; // USDC amount as number or null
  image: string;
  categoryId?: number | null; // Optional category ID
  donors?: string[]; // Optional for backward compatibility (was 'Donators')
  donations?: number[]; // Optional for backward compatibility
  withdrawn?: boolean; // Optional for backward compatibility
}

export interface UnsignedTransaction {
  to: string; // Contract address
  data: string; // Transaction data (hex)
  value: string; // Value in USDC (hex)
  gas: string; // Gas limit (hex)
  creationId?: string; // Campaign creation identifier (only for campaign creation)
}