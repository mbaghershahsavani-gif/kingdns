#!/bin/bash
set -e

echo "=== KingDNS Installer ==="

ROLE=${1:-international}
REGION=${2:-unknown}

echo "Installing KingDNS node"
echo "Role: $ROLE"
echo "Region: $REGION"

sudo apt update
sudo apt install -y git golang docker.io

sudo systemctl enable docker
sudo systemctl start docker

if [ ! -d /opt/kingdns ]; then
  sudo git clone https://github.com/mbaghershahsavani-gif/kingdns.git /opt/kingdns
fi

cd /opt/kingdns/dns-engine

go mod tidy
go fmt ./...
go build ./...

sudo mkdir -p /etc/kingdns

cat <<EOF | sudo tee /etc/kingdns/node.yaml
node:
  role: $ROLE
  region: $REGION
EOF

echo "KingDNS installation completed"
echo "Start runtime with:"
echo "cd /opt/kingdns/dns-engine && go run ./cmd/server"
