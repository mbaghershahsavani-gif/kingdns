package resolver

import (
	"strings"

	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/database"
	"github.com/miekg/dns"
)

var domainQuery database.DomainQuery

func ControllerHandler(w dns.ResponseWriter, req *dns.Msg) {
	msg := new(dns.Msg)
	msg.SetReply(req)

	if len(req.Question) > 0 {
		name := strings.TrimSuffix(req.Question[0].Name, ".")

		record := domainQuery.Find(name)

		if record.Value != "" {
			AddARecord(msg, req.Question[0].Name, record.Value, record.TTL)
		}
	}

	_ = w.WriteMsg(msg)
}
