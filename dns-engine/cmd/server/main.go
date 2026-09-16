package main

import (
	"log"
	"os"
	"sync"

	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/internal/cluster"
	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/internal/health"
	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/resolver"
	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/version"
	"github.com/miekg/dns"
)

func main() {
	log.Printf(
		"%s %s | Mode: %s",
		version.Name,
		version.Version,
		version.Mode,
	)

	region := os.Getenv("KINGDNS_REGION")

	if region == "" {
		region = "unknown"
	}

	log.Printf("Region: %s", region)

	go func() {
		health.StartServer(":8080")
	}()

	cluster.Bootstrap()

	go func() {
		cluster.StartServer(":8081")
	}()

	go func() {
		cluster.StartMonitor()
	}()

	go func() {
		cluster.StartSyncMonitor()
	}()

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
