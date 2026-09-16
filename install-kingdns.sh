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
echo "Building KingDNS binary"

mkdir -p /opt/kingdns/bin

cd /opt/kingdns/dns-engine

go mod tidy

go build -o /opt/kingdns/bin/kingdns ./cmd/server


echo "Installing systemd service"

cat > /etc/systemd/system/kingdns.service <<EOF
[Unit]
Description=KingDNS DNS Engine
After=network-online.target
Wants=network-online.target

[Service]
Type=simple

WorkingDirectory=/opt/kingdns/dns-engine

ExecStart=/opt/kingdns/bin/kingdns

Restart=always
RestartSec=5

Environment=KINGDNS_REGION=$REGION

StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target
EOF


systemctl daemon-reload

systemctl enable kingdns

systemctl restart kingdns

echo ""
echo "Installing KingDNS system service"

cp /opt/kingdns/dns-engine/deployment/systemd/kingdns.service /etc/systemd/system/kingdns.service

systemctl daemon-reload

systemctl enable kingdns

systemctl restart kingdns

echo ""
echo "KingDNS service:"
systemctl status kingdns --no-pager