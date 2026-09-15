CREATE TABLE IF NOT EXISTS routing_rules (
    id SERIAL PRIMARY KEY,
    service_id INTEGER REFERENCES services(id),
    condition TEXT NOT NULL,
    target_node INTEGER REFERENCES nodes(id),
    priority INTEGER DEFAULT 100,
    enabled BOOLEAN DEFAULT TRUE
);
