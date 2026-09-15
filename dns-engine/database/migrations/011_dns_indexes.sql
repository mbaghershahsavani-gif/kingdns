-- v2.9 resolver optimization

CREATE INDEX IF NOT EXISTS idx_domains_name
ON domains(name);
