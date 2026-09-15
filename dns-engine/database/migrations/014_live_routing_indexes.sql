CREATE TABLE IF NOT EXISTS live_routes (
	id SERIAL PRIMARY KEY,
	domain TEXT NOT NULL,
	target TEXT NOT NULL,
	enabled BOOLEAN DEFAULT true
);

CREATE INDEX IF NOT EXISTS idx_live_routes_domain
ON live_routes(domain);
