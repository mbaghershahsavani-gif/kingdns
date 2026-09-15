CREATE TABLE IF NOT EXISTS routing_rules (
	id SERIAL PRIMARY KEY,
	domain TEXT NOT NULL,
	target TEXT NOT NULL,
	enabled BOOLEAN DEFAULT true
);

CREATE INDEX IF NOT EXISTS idx_routing_rules_domain
ON routing_rules(domain);
