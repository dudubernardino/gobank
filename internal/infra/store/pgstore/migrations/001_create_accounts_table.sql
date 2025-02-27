-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS accounts (
  id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  tax_id VARCHAR(20) NOT NULL UNIQUE,
  name VARCHAR(255) NOT NULL,
  monthly_income BIGINT,
  annual_revenue BIGINT, 
  email VARCHAR(255) NOT NULL UNIQUE,
  balance BIGINT NOT NULL DEFAULT 0, 
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

---- create above / drop below ----

DROP TABLE IF EXISTS accounts;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
