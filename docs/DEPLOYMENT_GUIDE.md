# Production Deployment Guide

Recommended deployment:

Server 1:
Iran DNS Node

Server 2:
International DNS Node

Requirements:

- Ubuntu 22.04+
- Public IP
- UDP/TCP port 53
- Go runtime
- Docker support

Installation:

git clone https://github.com/mbaghershahsavani-gif/kingdns.git

cd kingdns/dns-engine

go build ./...

Start:

go run ./cmd/server

Validate:

nslookup kingdns.local SERVER_IP
