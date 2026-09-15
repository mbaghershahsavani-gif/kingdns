package resolver

import (
	"strings"

	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/database"
	"github.com/miekg/dns"
)

var repository = database.DomainRepository{}

func RuntimeHandler(w dns.ResponseWriter, req *dns.Msg) {
	msg := new(dns.Msg)
	msg.SetReply(req)

	if len(req.Question) > 0 {
		name := strings.TrimSuffix(req.Question[0].Name, ".")

		record := repository.Find(name)

		if record.Value != "" && record.Type == "A" {
			AddARecord(msg, req.Question[0].Name, record.Value, record.TTL)
		}
	}

	_ = w.WriteMsg(msg)
}
