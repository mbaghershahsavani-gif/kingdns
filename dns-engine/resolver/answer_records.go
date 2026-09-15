package resolver

import "github.com/miekg/dns"

func AddAAnswer(msg *dns.Msg, name string, ip string) {
	// Add real A record response here.
}

func AddAAAAAnswer(msg *dns.Msg, name string, ip string) {
	// Add real AAAA record response here.
}

func AddCNAMEAnswer(msg *dns.Msg, name string, target string) {
	// Add real CNAME response here.
}
