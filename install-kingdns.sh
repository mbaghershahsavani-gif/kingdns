#!/bin/bash
set -e

if [ -d "/opt/kingdns" ]; then
    rm -rf /opt/kingdns
fi

git clone https://github.com/mbaghershahsavani-gif/kingdns.git /opt/kingdns

echo "Building KingDNS binary"

mkdir -p /opt/kingdns/bin

cd /opt/kingdns/dns-engine

go mod tidy

go build -o /opt/kingdns/bin/kingdns ./cmd/server

if [ ! -f "/opt/kingdns/bin/kingdns" ]; then
    echo "KingDNS binary build failed"
    exit 1
fi


echo "Installing KingDNS system service"

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


sleep 3

echo ""
echo "KingDNS runtime validation"

systemctl is-active --quiet kingdns && echo "✓ Service running"

ss -lntup | grep :53

echo ""
echo "KingDNS installation completed"