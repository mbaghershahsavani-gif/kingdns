# KingDNS DNS Engine v3.3

Real Controller API Integration Foundation.

Added:

- controller HTTP transport
- node API client foundation
- route API foundation
- live node selection
- routing repository preparation

Architecture:

DNS Query
 |
DNS Engine
 |
Authenticated Controller API
 |
Nodes + Routes
 |
Selected IP
 |
DNS Answer

Next:

- real JWT requests
- PostgreSQL routing_rules execution
- Redis cache
- heartbeat driven failover
