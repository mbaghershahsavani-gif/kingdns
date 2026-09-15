package resolver

import (
	"log"
)

func StartUDP(address string) {
	log.Println("KingDNS UDP DNS listener foundation:", address)

	// Future:
	// - bind UDP :53
	// - parse DNS packets
	// - resolve queries
	// - write DNS responses
}
