CREATE TABLE IF NOT EXISTS domains (
    id SERIAL PRIMARY KEY,
    service_id INTEGER REFERENCES services(id),
    domain_pattern TEXT NOT NULL
);
