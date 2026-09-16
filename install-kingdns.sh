#!/bin/bash
set -e

ROLE=${1:-dns-edge}
REGION=${2:-international}

echo "================================"
echo " KingDNS Installer"
echo " Role: $ROLE"
echo " Region: $REGION"
echo "================================"

apt update

apt install -y git golang docker.io ufw

systemctl enable docker
systemctl start docker

if [ ! -d /opt/kingdns ]; then
    git clone https://github.com/mbaghershahsavani-gif/kingdns.git /opt/kingdns
else
    cd /opt/kingdns
    git pull
fi

cd /opt/kingdns/dns-engine

go mod tidy

go build ./...

mkdir -p /etc/kingdns

cat > /etc/kingdns/node.yaml <<EOF
node:
  role: $ROLE
  region: $REGION
EOF

echo ""
echo "KingDNS installation completed"
echo ""
echo "Configuration:"
cat /etc/kingdns/node.yaml
echo ""
echo "Installing KingDNS system service"

cp /opt/kingdns/dns-engine/deployment/systemd/kingdns.service /etc/systemd/system/kingdns.service

systemctl daemon-reload

systemctl enable kingdns

systemctl restart kingdns

echo ""
echo "KingDNS service:"
systemctl status kingdns --no-pager