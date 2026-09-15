# KingDNS DNS Engine v3.0

Controller Managed DNS Routing Foundation.

Implemented:

- routing engine foundation
- node selection policy
- health-aware routing structure
- cache integration point
- DNS query logging foundation
- database routing indexes

Architecture:

DNS Query
 |
Resolver
 |
Controller Routing
 |
Healthy Node Selection
 |
DNS Response

Next:

- real PostgreSQL queries
- Redis cache execution
- heartbeat driven routing
- geo and latency routing
