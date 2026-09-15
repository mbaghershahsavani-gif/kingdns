# KingDNS DNS Engine v3.1

Controller Communication Layer.

Added:

- controller API client foundation
- node discovery layer
- route client foundation
- controller-based node selection
- heartbeat sync preparation

Architecture:

DNS Query
 |
DNS Engine
 |
Controller API
 |
Nodes / Routes
 |
Selected Node
 |
DNS Response

Next:

- real HTTP controller calls
- authentication token
- heartbeat-driven routing
- Redis cache
- PostgreSQL routing rules
