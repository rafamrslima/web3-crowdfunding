import { type FormEvent, useState } from "react";
import { parseEther } from "viem"; // safe ETH -> wei conversion
import { API_BASE } from "./config";

function toUnixSeconds(dateStr: string): string {
  // dateStr expected from <input type="date"> or datetime-local (YYYY-MM-DD or ISO)
  const ms = new Date(dateStr).getTime();
  return Math.floor(ms / 1000).toString();
}

export default function App() {
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

  async function onSubmit(e: FormEvent) {
    e.preventDefault();

    // Basic guard: deadline must be in the future
    const nowSec = Math.floor(Date.now() / 1000);
    const deadlineSec = Number(toUnixSeconds(deadlineDate));
    if (deadlineSec <= nowSec) {
      alert("Deadline must be in the future.");
      return;
    }

    // Convert ETH → wei (as string)
    let targetWeiStr = "0";
    try {
      const wei = parseEther(targetEth); // returns bigint wei
      targetWeiStr = wei.toString();
    } catch {
      alert("Invalid ETH amount");
      return;
    }

    const payload = {
      owner: owner,
      title: title,
      description: description,
      target: targetWeiStr,               // wei (string)
      deadline: deadlineSec.toString(),   // unix seconds (string)
      image: image || ""                  // keep as empty string if not provided
    };

    try {
      const res = await fetch(`${API_BASE}/campaign/create-unsigned`, {
        method: "POST",
        headers: { "content-type": "application/json" },
        body: JSON.stringify(payload)
      });
      const data = await res.json().catch(() => ({}));
      console.log("API response:", data); // log payload body
    if (!res.ok) {
      console.error("API error:", res.status, res.statusText);
    } else {
      console.log("Request sent successfully");
    }
    } catch (err: unknown) {
      console.error(err);
      alert("Network error. See console for details.");
    }
  }

  // set min date = today
  const todayISO = new Date().toISOString().slice(0, 10);

  return (
    <div style={{ maxWidth: 640, margin: "2rem auto", fontFamily: "system-ui, sans-serif" }}>
      <h2>Create Campaign</h2>

      <form onSubmit={onSubmit} style={{ display: "grid", gap: 12 }}>
        <label>
          Owner (0x address)
          <input
            required
            pattern="^0x[a-fA-F0-9]{40}$"
            value={owner}
            onChange={(e) => setOwner(e.target.value)}
            placeholder="0x..."
            style={{ width: "100%" }}
          />
        </label>

        <label>
          Title
          <input
            required
            minLength={3}
            maxLength={80}
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            placeholder="Save the Amazon river"
            style={{ width: "100%" }}
          />
        </label>

        <label>
          Description
          <textarea
            required
            minLength={10}
            maxLength={1000}
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            placeholder="Funding reforestation projects worldwide."
            rows={4}
            style={{ width: "100%" }}
          />
        </label>

        <label>
          Target (in ETH)
          <input
            required
            inputMode="decimal"
            min="0"
            step="0.000000000000000001"
            value={targetEth}
            onChange={(e) => setTargetEth(e.target.value)}
            placeholder="234"
            style={{ width: "100%" }}
          />
        </label>

        <label>
          Deadline (date in the future)
          <input
            required
            type="date"
            min={todayISO}
            value={deadlineDate}
            onChange={(e) => setDeadlineDate(e.target.value)}
            style={{ width: "100%" }}
          />
        </label>

        <label>
          Image URL (optional)
          <input
            type="url"
            value={image}
            onChange={(e) => setImage(e.target.value)}
            placeholder="https://example.com/image.png"
            style={{ width: "100%" }}
          />
        </label>

        <button type="submit">Create campaign</button>
      </form>
    </div>
  );
}