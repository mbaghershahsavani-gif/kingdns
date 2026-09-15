package resolver

import "github.com/miekg/dns"

func BuildDNSResponse(req *dns.Msg, record DNSRecord) *dns.Msg {
	msg := new(dns.Msg)
	msg.SetReply(req)

	// Production answer generation foundation.
	// Future:
	// - A records
	// - AAAA records
	// - CNAME records

	return msg
}
