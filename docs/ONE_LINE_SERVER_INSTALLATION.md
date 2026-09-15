# KingDNS One Line Server Installation

## Iran Server

Run:

```bash
curl -fsSL https://raw.githubusercontent.com/mbaghershahsavani-gif/kingdns/main/install-kingdns.sh | bash -s dns-edge iran
```

## International Server

Run:

```bash
curl -fsSL https://raw.githubusercontent.com/mbaghershahsavani-gif/kingdns/main/install-kingdns.sh | bash -s dns-edge international
```

The same installer is used on both servers.

Configuration is generated automatically:

Iran:

role: dns-edge
region: iran

International:

role: dns-edge
region: international

After installation:

1. Start KingDNS
2. Verify health
3. Connect nodes to control plane
4. Test failover

Production improvements to add:

- systemd service
- TLS certificates
- automatic upgrades
- monitoring agent
