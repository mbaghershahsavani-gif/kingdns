# KingDNS DNS Engine v3.6

Real Data Plane Connections Foundation.

Added:

- PostgreSQL live client foundation
- Redis live cache foundation
- authenticated controller client foundation
- live routing execution path
- data plane resolver integration

Flow:

DNS Query
 |
Redis Cache
 |
PostgreSQL Routing
 |
Controller Node Health
 |
DNS Response

Next:

- production database pooling
- real Redis connection
- controller JWT verification
- multi-node failover
- geo latency routing
