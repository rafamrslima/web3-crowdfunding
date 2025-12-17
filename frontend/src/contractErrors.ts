// Contract error handling utilities
import { BaseError, ContractFunctionRevertedError } from 'viem';

/**
 * Extract raw contract error message without any mapping or friendly messages
 */
export function getRawContractError(error: unknown): string {
  if (error instanceof BaseError) {
    const revertError = error.walk(err => err instanceof ContractFunctionRevertedError);
    
    if (revertError instanceof ContractFunctionRevertedError) {
      // Return the raw revert reason string (require message)
      if (revertError.reason) {
        return revertError.reason;
      }
    }

    // Try to extract revert reason from error message
    if (error.message?.includes('execution reverted')) {
      const revertMatch = error.message.match(/execution reverted:?\s*(.+?)(?:\s*$|\s*\(|$)/i);
      if (revertMatch && revertMatch[1]) {
        return revertMatch[1].trim().replace(/"/g, '');
      }
    }

    // Return the original error message as-is
    return error.message || error.shortMessage || 'Transaction failed';
  }

  // Fallback for non-viem errors - return original message
  if (error instanceof Error) {
    return error.message;
  }

  return 'Unknown error occurred';
}