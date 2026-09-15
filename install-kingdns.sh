#!/bin/bash
set -e

ROLE=${1:-dns-edge}
REGION=${2:-international}

echo "Installing KingDNS production bootstrap"
echo "Role: $ROLE"
echo "Region: $REGION"

sudo apt update
sudo apt install -y git golang docker.io ufw

sudo systemctl enable docker
sudo systemctl start docker

if [ ! -d /opt/kingdns ]; then
  sudo git clone https://github.com/mbaghershahsavani-gif/kingdns.git /opt/kingdns
else
  cd /opt/kingdns && sudo git pull
fi

cd /opt/kingdns/dns-engine

go mod tidy
go build ./...

sudo mkdir -p /etc/kingdns

cat <<EOF | sudo tee /etc/kingdns/node.yaml
node:
  role: $ROLE
  region: $REGION

security:
  tls: enabled

monitoring:
  health: enabled
EOF

sudo cp deployment/systemd/kingdns.service /etc/systemd/system/kingdns.service

sudo systemctl daemon-reload
sudo systemctl enable kingdns
sudo systemctl restart kingdns

echo "KingDNS bootstrap completed"
