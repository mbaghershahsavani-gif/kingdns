# KingDNS Architecture

## Overview

KingDNS is built as a distributed DNS platform.

Layers:

1. Resolver Engine
2. Routing Engine
3. Control Plane
4. Security Layer
5. Intelligence Layer
6. Deployment Layer

## Multi Region Model

Iran Node:

- Primary regional DNS
- Local routing
- Health reporting

International Node:

- Failover DNS
- Global routing
- Disaster recovery

Architecture:

Users
 |
Control Plane
 |
+-------------+
|             |
Iran       International
Node          Node
