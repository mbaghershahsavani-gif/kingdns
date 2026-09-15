package resolver

import (
	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/routing"
	"github.com/miekg/dns"
)

func DataPlaneHandler(w dns.ResponseWriter, req *dns.Msg) {
	msg := new(dns.Msg)
	msg.SetReply(req)

	if len(req.Question) > 0 {
		target := routing.Resolve(req.Question[0].Name)

		if target != "" {
			AddARecord(msg, req.Question[0].Name, target, 60)
		}
	}

	_ = w.WriteMsg(msg)
}
