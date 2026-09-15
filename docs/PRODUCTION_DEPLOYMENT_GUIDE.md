\# KingDNS Production Deployment Guide



\## Overview



This guide explains how to deploy KingDNS on real servers.



Recommended first production test topology:



```

&#x20;                Users



&#x20;                  |



&#x20;            KingDNS Control Plane



&#x20;             /                 \\



&#x20;            /                   \\



&#x20;   Iran DNS Node          International DNS Node



&#x20;     Server 1                 Server 2



&#x20;   Primary Region          Failover Region

```



The recommended deployment:



\* Server 1: Iran

\* Server 2: International location



Recommended international locations:



\* Germany

\* Turkey

\* Netherlands

\* UAE

\* Singapore



\---



\# Server Requirements



\## Minimum Test Environment



Each server:



\* Ubuntu 22.04 LTS or newer

\* 2 CPU cores

\* 2-4 GB RAM

\* 40 GB SSD

\* Public IPv4 address

\* Open DNS ports



Required ports:



```

UDP 53

TCP 53

TCP 8080 (health API)

TCP 443 (secure communication)

```



\---



\# Step 1 - Prepare Servers



Update both servers:



```bash

sudo apt update

sudo apt upgrade -y

```



Install required packages:



```bash

sudo apt install -y git golang docker.io

```



Enable Docker:



```bash

sudo systemctl enable docker

sudo systemctl start docker

```



\---



\# Step 2 - Download KingDNS



On each server:



```bash

git clone https://github.com/mbaghershahsavani-gif/kingdns.git



cd kingdns/dns-engine

```



\---



\# Step 3 - Configure Iran Node



Create:



```

config/node.yaml

```



Example:



```yaml

node:

&#x20; id: iran-01

&#x20; region: iran

&#x20; role: dns-edge



controller:

&#x20; enabled: true



routing:

&#x20; mode: intelligent

```



This server acts as the primary regional DNS node.



\---



\# Step 4 - Configure International Node



Example:



```yaml

node:

&#x20; id: international-01

&#x20; region: international

&#x20; role: dns-edge



controller:

&#x20; enabled: true



routing:

&#x20; mode: intelligent

```



This server provides:



\* redundancy

\* failover

\* international routing



\---



\# Step 5 - Build KingDNS



On both servers:



```bash

go mod tidy



go fmt ./...



go build ./...

```



Verify:



```bash

go run ./cmd/server

```



\---



\# Step 6 - Verify DNS Service



From another machine:



```bash

nslookup kingdns.local SERVER\_IP

```



Expected:



```

Name: kingdns.local

Address: SERVER\_IP

```



\---



\# Step 7 - Enable Node Communication



Each node must report:



\* health status

\* latency

\* availability

\* routing information



Communication flow:



```

Iran Node

&#x20;    |

&#x20;    |

Heartbeat

&#x20;    |

&#x20;    |

Control Plane

&#x20;    |

&#x20;    |

Heartbeat

&#x20;    |

International Node

```



\---



\# Step 8 - Test Failover



Normal operation:



```

Client



&#x20;|



Iran Node



&#x20;|



DNS Response

```



Stop Iran DNS service:



```bash

sudo systemctl stop kingdns

```



Expected:



```

Client



&#x20;|



International Node



&#x20;|



DNS Response

```



Restore:



```bash

sudo systemctl start kingdns

```



\---



\# Step 9 - Production DNS Delegation



For a real domain:



Create DNS records:



```

ns1.example.com -> Iran Server IP



ns2.example.com -> International Server IP

```



Register nameservers:



```

ns1.example.com

ns2.example.com

```



Then test:



```bash

dig example.com NS

```



\---



\# Monitoring



Monitor:



\* DNS response time

\* node health

\* failed queries

\* traffic distribution

\* failover events



Recommended production stack:



\* Prometheus

\* Grafana

\* Loki

\* Alertmanager



\---



\# Security Recommendations



Before production:



Enable:



\* TLS node communication

\* firewall rules

\* SSH key authentication

\* automatic updates

\* backup strategy



Never expose administrative interfaces publicly without authentication.



\---



\# Disaster Recovery



Maintain:



\* configuration backups

\* database backups

\* node recovery procedures

\* DNS rollback plan



\---



\# Deployment Checklist



\## Server Preparation



\[ ] Ubuntu installed



\[ ] Firewall configured



\[ ] Docker installed



\[ ] Go installed



\## KingDNS Installation



\[ ] Repository cloned



\[ ] Configuration created



\[ ] Build successful



\[ ] DNS service running



\## Multi Region



\[ ] Iran node connected



\[ ] International node connected



\[ ] Heartbeat working



\[ ] Failover tested



\## Production



\[ ] DNS delegation completed



\[ ] Monitoring enabled



\[ ] Backup enabled



\---



\# Support



The KingDNS architecture supports:



\* multi-region DNS

\* intelligent routing

\* health-based failover

\* enterprise security

\* multi-tenant operation

\* global deployment



This guide represents the recommended production deployment model.



