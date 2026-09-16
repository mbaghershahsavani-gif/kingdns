echo "Building KingDNS binary"

mkdir -p /opt/kingdns/bin

cd /opt/kingdns/dns-engine

go mod tidy

go build -o /opt/kingdns/bin/kingdns ./cmd/server


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


echo ""
echo "KingDNS service status:"
systemctl status kingdns --no-pager