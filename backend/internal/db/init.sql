CREATE TABLE IF NOT EXISTS campaigns (
  campaign_id    BIGINT PRIMARY KEY,
  owner          BYTEA NOT NULL,
  target_amount  NUMERIC(78,0) NOT NULL,
  deadline_ts    BIGINT NOT NULL,
  tx_hash        BYTEA NOT NULL,
  block_number   BIGINT NOT NULL,
  block_time     TIMESTAMPTZ,
  created_at     TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS donations (
  id             BIGSERIAL PRIMARY KEY,
  campaign_id    BIGINT,
  donor          BYTEA NOT NULL,
  amount         NUMERIC(78,0) NOT NULL,
  tx_hash        BYTEA NOT NULL,
  block_number   BIGINT NOT NULL,
  block_time     TIMESTAMPTZ,
  created_at     TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS refunds (
  campaign_id       BIGINT,
  donor             BYTEA NOT NULL,
  total_contributed NUMERIC(78,0) NOT NULL,
  tx_hash           BYTEA NOT NULL,
  block_number      BIGINT NOT NULL,
  block_time        TIMESTAMPTZ,
  created_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
  PRIMARY KEY (campaign_id, donor)
);

CREATE TABLE IF NOT EXISTS withdrawals (
  campaign_id    BIGINT,
  owner          BYTEA NOT NULL,
  amount         NUMERIC(78,0) NOT NULL,
  tx_hash        BYTEA NOT NULL,
  block_number   BIGINT NOT NULL,
  block_time     TIMESTAMPTZ,
  created_at     TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS sync_state (
  id         TEXT PRIMARY KEY,
  last_block BIGINT NOT NULL
);

INSERT INTO sync_state (id, last_block)
VALUES ('crowdfunding-local', :START_BLOCK - 1)
ON CONFLICT (id) DO NOTHING;