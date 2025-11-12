import { FormEvent, useState } from "react";
import { createWalletClient, custom, parseEther } from "viem";
import type { Hex } from "viem";
import { CHAIN_ID, API_BASE } from "./config";
import "./App.css";

declare global { interface Window { ethereum?: any } }

function toUnixSeconds(dateStr: string): string {
  const ms = new Date(dateStr).getTime();
  return Math.floor(ms / 1000).toString();
}

export default function App() {
  const [account, setAccount] = useState<`0x${string}` | null>(null);
  const walletClient = window.ethereum
    ? createWalletClient({ chain: { id: CHAIN_ID } as any, transport: custom(window.ethereum) })
    : null;
  // Defaults (editable in the UI)
  const [owner, setOwner] = useState("0x70997970C51812dc3A010C7d01b50e0d17dc79C8");
  const [title, setTitle] = useState("Save the Amazon river");
  const [description, setDescription] = useState("Funding reforestation projects worldwide.");
  // Show target in ETH to user; convert to wei before sending
  const [targetEth, setTargetEth] = useState("1"); // 1 ETH (→ 1e18 wei when sending)
  // Default deadline from the given Unix ts (1794268800) ≈ 2026-...; show as date input
  const defaultDeadlineISO = new Date(1794268800 * 1000).toISOString().slice(0, 10);
  const [deadlineDate, setDeadlineDate] = useState(defaultDeadlineISO);
  const [image, setImage] = useState("");

  // UI state management
  const [isCreating, setIsCreating] = useState(false);
  const [successTxHash, setSuccessTxHash] = useState<string | null>(null);
  const [errorMessage, setErrorMessage] = useState<string | null>(null);

  // Clear messages when user starts editing form
  const clearMessages = () => {
    if (successTxHash || errorMessage) {
      setSuccessTxHash(null);
      setErrorMessage(null);
    }
  };

  async function connectWallet() {
    if (!walletClient) {
      console.error("MetaMask not found");
      return;
    }
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
    setOwner(addr); // preload owner field, but it stays editable
  }

  function disconnectWallet() {
    setAccount(null);
    setOwner("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"); // Reset to default
    setSuccessTxHash(null);
    setErrorMessage(null);
  }

  async function onSubmit(e: FormEvent) {
    e.preventDefault();

    // Reset previous states
    setErrorMessage(null);
    setSuccessTxHash(null);
    setIsCreating(true);

    try {
      if (!walletClient || !account) {
        throw new Error("Please connect wallet first");
      }

      // Basic guard: deadline must be in the future
      const nowSec = Math.floor(Date.now() / 1000);
      const deadlineSec = Number(toUnixSeconds(deadlineDate));
      if (deadlineSec <= nowSec) {
        throw new Error("Deadline must be in the future");
      }

      // Convert ETH → wei (as string)
      let targetWeiStr = "0";
      try {
        const wei = parseEther(targetEth); // returns bigint wei
        targetWeiStr = wei.toString();
      } catch {
        throw new Error("Invalid ETH amount");
      }

      const payload = {
        owner: owner,
        title: title,
        description: description,
        target: targetWeiStr,               // wei (string)
        deadline: deadlineSec.toString(),   // unix seconds (string)
        image: image || ""                  // keep as empty string if not provided
      };

      console.log("Creating unsigned transaction...");
      const res = await fetch(`${API_BASE}/campaign/create-unsigned`, {
        method: "POST",
        headers: { "content-type": "application/json" },
        body: JSON.stringify(payload)
      });

      const built = await res.json();
      console.log("API response (unsigned tx):", built);
      
      if (!res.ok) {
        throw new Error(`API error: ${res.status} ${res.statusText} - ${JSON.stringify(built)}`);
      }

      console.log("Requesting MetaMask signature...");
      const txHash = await walletClient.sendTransaction({
        account,
        to: built.to as `0x${string}`,
        data: built.data as Hex,
        // value/gas come as hex strings; convert to BigInt if present
        value: built.value ? BigInt(built.value) : 0n,
        gas: built.gas ? BigInt(built.gas) : undefined,
        chain: undefined
        // omit fees/nonce; wallet will populate
      });
      
      console.log("✅ Transaction successful:", txHash);
      setSuccessTxHash(txHash);
      
    } catch (err) {
      console.error("Campaign creation error:", err);
      const error = err as Error;
      
      // Handle specific error types
      if (error.message?.includes("User rejected") || error.message?.includes("rejected")) {
        setErrorMessage("Transaction was rejected by user");
      } else if (error.message?.includes("insufficient funds")) {
        setErrorMessage("Insufficient ETH balance for transaction");
      } else {
        setErrorMessage(error.message || "Failed to create campaign");
      }
    } finally {
      setIsCreating(false);
    }
  }

  // set min date = today
  const todayISO = new Date().toISOString().slice(0, 10);

  return (
    <div className="app-container">
      <h1 className="app-title">🚀 Create Crowdfunding Campaign</h1>

      <div className="wallet-section">
        <h3>1. Connect Your Wallet</h3>
        <div className="wallet-controls">
          {!account ? (
            <button onClick={connectWallet} className="btn btn-primary">
              🦊 Connect MetaMask
            </button>
          ) : (
            <>
              <button className="btn btn-connected" disabled>
                ✅ <span className="wallet-address">{account.slice(0, 6)}…{account.slice(-4)}</span>
              </button>
              <button onClick={disconnectWallet} className="btn btn-danger">
                🔌 Disconnect
              </button>
            </>
          )}
        </div>
      </div>

      <div className="form-container">
        <h3>2. Campaign Details</h3>
        <form onSubmit={onSubmit} className="form">
          <div className="form-group">
            <label className="form-label">Owner (0x address)</label>
            <input
              required
              pattern="^0x[a-fA-F0-9]{40}$"
              value={owner}
              onChange={(e) => setOwner(e.target.value)}
              placeholder="0x..."
              className="form-input"
            />
          </div>

          <div className="form-group">
            <label className="form-label">Title</label>
            <input
              required
              minLength={3} maxLength={80}
              value={title}
              onChange={(e) => {
                setTitle(e.target.value);
                clearMessages();
              }}
              placeholder="Save the Amazon river"
              className="form-input"
            />
          </div>

          <div className="form-group">
            <label className="form-label">Description</label>
            <textarea
              required
              minLength={10} maxLength={1000}
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              placeholder="Funding reforestation projects worldwide."
              rows={4}
              className="form-input form-textarea"
            />
          </div>

          <div className="form-group">
            <label className="form-label">Target (in ETH)</label>
            <input
              required
              inputMode="decimal"
              min="0"
              step="0.000000000000000001"
              value={targetEth}
              onChange={(e) => setTargetEth(e.target.value)}
              placeholder="1"
              className="form-input"
            />
          </div>

          <div className="form-group">
            <label className="form-label">Deadline (date in the future)</label>
            <input
              required
              type="date"
              min={todayISO}
              value={deadlineDate}
              onChange={(e) => setDeadlineDate(e.target.value)}
              className="form-input"
            />
          </div>

          <div className="form-group">
            <label className="form-label">Image URL (optional)</label>
            <input
              type="url"
              value={image}
              onChange={(e) => setImage(e.target.value)}
              placeholder="https://example.com/image.png"
              className="form-input"
            />
          </div>

          <button 
            type="submit" 
            disabled={isCreating || !account}
            className={`btn submit-button ${isCreating ? 'btn-secondary' : 'btn-success'}`}
          >
            {isCreating && <span className="loading-spinner"></span>}
            {isCreating ? 'Creating Campaign...' : 'Create Campaign'}
          </button>

          {/* Success Message */}
          {successTxHash && (
            <div className="message-box message-success">
              <h4 className="message-title">✅ Campaign Created Successfully!</h4>
              <p className="message-text">Your campaign has been created and submitted to the blockchain.</p>
              <p className="message-text">
                <strong>Transaction ID:</strong>
                <code className="transaction-hash">{successTxHash}</code>
              </p>
            </div>
          )}

          {/* Error Message */}
          {errorMessage && (
            <div className="message-box message-error">
              <h4 className="message-title">❌ Error</h4>
              <p className="message-text">{errorMessage}</p>
            </div>
          )}
        </form>
      </div>
    </div>
  );
}