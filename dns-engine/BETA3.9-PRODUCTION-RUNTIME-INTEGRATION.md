# KingDNS DNS Engine v3.9

Production Runtime Integration Foundation.

Added:

- PostgreSQL runtime lookup path
- Redis runtime cache path
- controller heartbeat foundation
- node registry foundation
- runtime routing selector

Architecture:

DNS Query
 |
Redis Runtime Cache
 |
PostgreSQL Runtime Lookup
 |
Controller Heartbeat
 |
Node Registry
 |
Runtime Routing
 |
DNS Response

Next:

- real database drivers
- real Redis connection
- controller API transport
- distributed node health
- production deployment
