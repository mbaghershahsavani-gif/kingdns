package resolver

import "log"

func StartResolver() {
	log.Println("KingDNS Resolver v1.5 runtime foundation")

	// Production next:
	// - bind UDP :53
	// - bind TCP :53
	// - process DNS packets
}
