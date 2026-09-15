# KingDNS DNS Engine v4.0

Distributed DNS Control Plane Foundation.

Added:

- PostgreSQL pool architecture
- distributed cache architecture
- controller control-plane foundation
- heartbeat scheduling foundation
- distributed health model
- multi-node routing foundation
- deployment structure

Architecture:

DNS Query
 |
Distributed Cache
 |
Database Pool
 |
Controller Sync
 |
Node Health
 |
Distributed Routing
 |
DNS Response

Next:

- production pgx driver
- real Redis cluster
- controller API transport
- multi-region routing
- production monitoring
