package resolver

import (
	"github.com/miekg/dns"
)

func LiveRoutingHandler(w dns.ResponseWriter, req *dns.Msg) {
	msg := new(dns.Msg)
	msg.SetReply(req)

	if len(req.Question) > 0 {
		AddARecord(msg, req.Question[0].Name, "127.0.0.1", 60)
	}

	_ = w.WriteMsg(msg)
}
