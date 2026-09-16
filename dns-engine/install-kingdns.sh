echo "Installing systemd service"

cp deployment/systemd/kingdns.service /etc/systemd/system/kingdns.service

systemctl daemon-reload

systemctl enable kingdns

systemctl restart kingdns

echo ""
echo "KingDNS service status:"
systemctl status kingdns --no-pager