package resolver

import "github.com/miekg/dns"

func HandleDNS(w dns.ResponseWriter, r *dns.Msg) {
	// Future:
	// - inspect question
	// - resolve domain
	// - create DNS answer
}
