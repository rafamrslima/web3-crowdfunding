import { parseUnits, getContract } from 'viem';
import type { Hex, WalletClient } from 'viem';
import { USDC_TOKEN_ADDRESS, CROWDFUNDING_CONTRACT_ADDRESS } from '../config';
import MockUSDCAbi from '../MockUSDC.abi.json';

// Get USDC contract instance
export const getUSDCContract = (walletClient: WalletClient) => {
  return getContract({
    address: USDC_TOKEN_ADDRESS as Hex,
    abi: MockUSDCAbi,
    client: walletClient
  });
};

// Get user's USDC balance
export const getUSDCBalance = async (walletClient: WalletClient, userAddress: string): Promise<string> => {
  const contract = getUSDCContract(walletClient);
  try {
    const balance = await contract.read.balanceOf([userAddress as Hex]) as bigint;
    // Convert from micro-units to dollars (divide by 1,000,000)
    const balanceInDollars = Number(balance) / 1000000;
    return balanceInDollars.toFixed(2);
  } catch (error) {
    console.error('Error getting USDC balance:', error);
    return '0';
  }
};

// Check current allowance
export const getUSDCAllowance = async (walletClient: WalletClient, owner: string): Promise<bigint> => {
  const contract = getUSDCContract(walletClient);
  try {
    const allowance = await contract.read.allowance([
      owner as Hex, 
      CROWDFUNDING_CONTRACT_ADDRESS as Hex
    ]) as bigint;
    return allowance;
  } catch (error) {
    console.error('Error getting USDC allowance:', error);
    return 0n;
  }
};

// Approve USDC spending for crowdfunding contract
export const approveUSDC = async (
  walletClient: WalletClient,
  account: Hex,
  amountInDollars: string // e.g., "10.5"
): Promise<Hex> => {
  
  // Convert dollars to micro-units (multiply by 1,000,000)
  const amountRaw = parseUnits(amountInDollars, 6);
  
  const contract = getUSDCContract(walletClient);
  
  try {
    // Execute approval transaction
    const hash = await contract.write.approve([
      CROWDFUNDING_CONTRACT_ADDRESS as Hex,
      amountRaw
    ], {
      account
    });

    return hash;
  } catch (error) {
    console.error('Error approving USDC:', error);
    throw new Error('Failed to approve USDC spending. Please try again.');
  }
};

// Check if approval is needed (returns true if need approval)
export const needsApproval = async (
  walletClient: WalletClient,
  owner: string,
  amountInDollars: string
): Promise<boolean> => {
  try {
    const currentAllowance = await getUSDCAllowance(walletClient, owner);
    const requiredAmount = parseUnits(amountInDollars, 6);
    
    return currentAllowance < requiredAmount;
  } catch (error) {
    console.error('Error checking approval:', error);
    return true; // If we can't check, assume approval is needed
  }
};