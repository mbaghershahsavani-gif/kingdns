# KingDNS

## Enterprise Distributed DNS Platform

KingDNS is a distributed intelligent DNS platform designed for multi-region DNS operations.

Features:

- Distributed DNS runtime
- Intelligent routing
- Health-based failover
- Multi-region deployment
- Enterprise security foundation
- Multi-tenant DNS architecture
- Kubernetes/cloud deployment support

## Quick Deployment

### Iran Node 🇮🇷

```bash
curl -fsSL https://raw.githubusercontent.com/mbaghershahsavani-gif/kingdns/main/install-kingdns.sh | bash -s dns-edge iran
```

### International Node 🌍

```bash
curl -fsSL https://raw.githubusercontent.com/mbaghershahsavani-gif/kingdns/main/install-kingdns.sh | bash -s dns-edge international
```

## Architecture

KingDNS supports:

Client
 |
Global DNS Platform
 |
Control Plane
 |
Regional DNS Nodes
 |
Intelligent Routing

## Documentation

See:

- docs/ARCHITECTURE.md
- docs/ROADMAP.md
- docs/DEPLOYMENT_GUIDE.md
- docs/SECURITY.md

## License

Apache License 2.0
