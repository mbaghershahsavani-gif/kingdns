package resolver

import "github.com/miekg/dns"

func StartTCPServer(address string) error {
	server := &dns.Server{
		Addr: address,
		Net: "tcp",
	}

	return server.ListenAndServe()
}
