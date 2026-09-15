package main

import (
	"log"
	"sync"

	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/resolver"
	"github.com/miekg/dns"
)

func main() {
	log.Println("KingDNS DNS Engine v2.7 starting")

	// Register DNS request handler
	dns.HandleFunc(".", resolver.RuntimeHandler)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := resolver.StartUDPServer(":53"); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := resolver.StartTCPServer(":53"); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("DNS listeners started on UDP/TCP :53")

	wg.Wait()
}
