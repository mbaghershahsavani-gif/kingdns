# KingDNS DNS Engine v3.5 Routing Model Consolidation

Completed cleanup:

- single routing Node model
- single Decision model
- single Policy model
- removed duplicate Select functions
- separated routing decisions from DNS response data

Architecture:

DNS Query
 |
Resolver
 |
Routing Decision
 |
Node Selection
 |
DNS Response

Ready for v3.6 real PostgreSQL and Redis data plane integration.
