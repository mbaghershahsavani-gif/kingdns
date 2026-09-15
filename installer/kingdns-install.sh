#!/usr/bin/env bash
set -e

echo "Installing KingDNS..."

apt update
apt install -y curl nginx dnsmasq nftables

echo "KingDNS base installation completed."
