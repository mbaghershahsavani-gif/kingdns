CREATE TABLE IF NOT EXISTS nodes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(50) NOT NULL,
    country VARCHAR(10),
    ip_address VARCHAR(100),
    status VARCHAR(50) DEFAULT 'offline',
    last_seen TIMESTAMP
);
