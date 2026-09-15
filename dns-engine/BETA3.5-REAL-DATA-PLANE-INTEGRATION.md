# KingDNS DNS Engine v3.5

Real Data Plane Integration Foundation.

Added:

- PostgreSQL routing repository foundation
- Redis cache client foundation
- controller route execution layer
- data plane resolver pipeline
- routing table migration

Architecture:

DNS Query
 |
Resolver
 |
Redis Cache
 |
PostgreSQL Routing Rules
 |
Controller Decision
 |
DNS Response

Next:

- real PostgreSQL connection
- real Redis connection
- authenticated controller requests
- multi-node failover
- geo latency routing
