CREATE TABLE IF NOT EXISTS campaigns (
  campaignId     BIGINT PRIMARY KEY,
  owner          BYTEA NOT NULL,
  target_wei     NUMERIC(78,0) NOT NULL,
  deadline_ts    BIGINT NOT NULL,
  created_tx     BYTEA NOT NULL,
  created_block  BIGINT NOT NULL,
  created_at     TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS donations (
  id             BIGSERIAL PRIMARY KEY,
  campaignId     BIGINT,
  donor          BYTEA NOT NULL,
  amount_wei     NUMERIC(78,0) NOT NULL,
  tx_hash        BYTEA NOT NULL,
  created_block  BIGINT NOT NULL,
  created_at     TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS sync_state (
  id         TEXT PRIMARY KEY,
  last_block BIGINT NOT NULL
);

INSERT INTO sync_state (id, last_block)
VALUES ('crowdfunding-local', :START_BLOCK - 1)
ON CONFLICT (id) DO NOTHING;