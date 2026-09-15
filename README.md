# KingDNS

KingDNS is a next-generation smart DNS traffic management platform inspired by DoctorDNS.

## Architecture

* controller: Go control plane API
* dns-engine: routing and resolver intelligence
* agent: relay and exit node agents
* dashboard: Next.js administration panel
* installer: production installer scripts


# KingDNS DNS Engine



Enterprise distributed DNS platform with:



- intelligent routing

- multi-region DNS

- health-based failover

- enterprise security

- production deployment support


# Quick Production Deployment

KingDNS supports two-node deployment:

Iran DNS Node            International DNS Node

## Server 1 — Iran Node 🇮🇷

Run this command on the Iran server:

```bash

curl -fsSL https://raw.githubusercontent.com/mbaghershahsavani-gif/kingdns/main/install-kingdns.sh | bash -s dns-edge iran

```

Iran DNS Node            International DNS Node

## Server 2 — International Node

Run this command on the Iran server:

Server 2 — International Node 🌍

Run this command on the international server:

```bash

curl -fsSL https://raw.githubusercontent.com/mbaghershahsavani-gif/kingdns/main/install-kingdns.sh | bash -s dns-edge international

```

