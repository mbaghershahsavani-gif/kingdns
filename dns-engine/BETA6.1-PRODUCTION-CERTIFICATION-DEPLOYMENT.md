# KingDNS DNS Engine v6.1

Production Certification & Real Server Deployment Framework.

Deployment model:

Server 1:
Iran DNS Node
- primary regional node
- DNS runtime
- cache
- health reporting

Server 2:
International DNS Node
- secondary/failover node
- DNS runtime
- cache
- health reporting

Architecture:

Users
 |
KingDNS Control Plane
 |
+----------------+
|                |
Iran Node     International Node

Features:

- node registration
- heartbeat
- TLS identity foundation
- deployment validation
- regional configuration

Next:

- live VPS deployment
- real DNS delegation
- monitoring
- failover testing
