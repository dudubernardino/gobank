-- Write your migrate up statements here

ALTER TABLE IF EXISTS accounts ADD COLUMN deleted_at TIMESTAMPTZ;

---- create above / drop below ----

ALTER TABLE IF EXISTS accounts DROP COLUMN IF EXISTS deleted_at;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
