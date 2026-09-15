CREATE INDEX IF NOT EXISTS idx_routing_rules_domain
ON routing_rules(domain_id);

CREATE INDEX IF NOT EXISTS idx_nodes_status
ON nodes(status);
