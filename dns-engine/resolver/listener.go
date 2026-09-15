package resolver

import (
	"github.com/miekg/dns"
)

func StartUDP(address string) error {
	server := &dns.Server{
		Addr: address,
		Net: "udp",
	}

	return server.ListenAndServe()
}
