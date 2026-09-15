package resolver

import "github.com/miekg/dns"

func HandleQuery(w dns.ResponseWriter, req *dns.Msg) {
	msg := new(dns.Msg)
	msg.SetReply(req)

	// Future:
	// - inspect DNS question
	// - lookup records
	// - execute routing
	// - write answers

	_ = w.WriteMsg(msg)
}
