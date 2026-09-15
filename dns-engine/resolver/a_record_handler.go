package resolver

import (
	"net"

	"github.com/miekg/dns"
)

func AddARecord(msg *dns.Msg, name string, ip string, ttl uint32) {
	answer := &dns.A{
		Hdr: dns.RR_Header{
			Name:   name,
			Rrtype: dns.TypeA,
			Class:  dns.ClassINET,
			Ttl:    ttl,
		},
		A: net.ParseIP(ip),
	}

	msg.Answer = append(msg.Answer, answer)
}
