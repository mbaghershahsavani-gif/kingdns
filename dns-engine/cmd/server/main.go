package main

import (
	"log"

	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/resolver"
)

func main() {
	log.Println("KingDNS DNS Engine v1.3 starting")

	resolver.StartUDP(":53")
	resolver.StartTCP(":53")
}
