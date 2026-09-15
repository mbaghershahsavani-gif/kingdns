CREATE TABLE IF NOT EXISTS dns_queries (
    id BIGSERIAL PRIMARY KEY,
    domain TEXT NOT NULL,
    client_ip TEXT,
    node_id INTEGER REFERENCES nodes(id),
    latency_ms INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
