package resolver

import "github.com/miekg/dns"

func RuntimeHandler(w dns.ResponseWriter, req *dns.Msg) {
	msg := new(dns.Msg)
	msg.SetReply(req)

	// v2.6:
	// 1. read question
	// 2. lookup domain
	// 3. check cache
	// 4. select route
	// 5. build answer

	_ = w.WriteMsg(msg)
}
