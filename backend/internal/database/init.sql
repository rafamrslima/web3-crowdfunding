CREATE TABLE IF NOT EXISTS campaigns (
  campaign_id    BIGINT PRIMARY KEY,
  owner          BYTEA NOT NULL,
  title          VARCHAR(50),
  description    VARCHAR(300),
  target_amount  NUMERIC(78,0) NOT NULL,
  deadline_ts    BIGINT NOT NULL,
  image          VARCHAR(50),
  tx_hash        BYTEA NOT NULL,
  block_number   BIGINT NOT NULL,
  block_time     TIMESTAMPTZ,
  category_id    INTEGER,
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

CREATE TABLE IF NOT EXISTS campaign_drafts (
  creation_id    VARCHAR(66) NOT NULL UNIQUE,
  owner          BYTEA NOT NULL,
  title          VARCHAR(50),
  description    VARCHAR(300),
  image          VARCHAR(50),
  category_id    INTEGER,
  created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
  PRIMARY KEY (creation_id)
);

CREATE TABLE IF NOT EXISTS sync_state (
  chain_id             BIGINT PRIMARY KEY,
  last_processed_block BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS categories (
  id           SERIAL PRIMARY KEY,
  name         VARCHAR(50) NOT NULL UNIQUE,
  slug         VARCHAR(50) NOT NULL UNIQUE,
  description  VARCHAR(200),
  created_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);

INSERT INTO categories (name, slug, description) VALUES
  ('Environment', 'environment', 'Climate, conservation, and sustainability projects'),
  ('Technology', 'technology', 'Software, hardware, and innovation'),
  ('Health', 'health', 'Medical, wellness, and mental health initiatives'),
  ('Arts & Culture', 'arts-culture', 'Music, film, design, and literature'),
  ('Education', 'education', 'Schools, courses, and research'),
  ('Community', 'community', 'Local projects and social causes'),
  ('Sports', 'sports', 'Athletics, fitness, and teams'),
  ('Animals', 'animals', 'Wildlife, pets, and conservation'),
  ('Business', 'business', 'Startups, products, and services'),
  ('Gaming', 'gaming', 'Games, esports, and streaming'),
  ('Others', 'others', 'General projects and miscellaneous campaigns')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO sync_state (chain_id, last_processed_block)
VALUES ('31337', 0)
ON CONFLICT (chain_id) DO NOTHING;