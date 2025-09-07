CREATE TABLE IF NOT EXISTS items (
	id BIGSERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT NOT NULL DEFAULT ''
);

-- Useful index for ordering
CREATE INDEX IF NOT EXISTS idx_items_id_desc ON items (id DESC);
