package resolver

import (
	"log"
	"strings"

	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/database"
	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/internal/cluster"
	"github.com/miekg/dns"
)

var repository = database.DomainRepository{}

func RuntimeHandler(w dns.ResponseWriter, req *dns.Msg) {

	route := cluster.CurrentRoute()

	log.Printf(
		"DNS request routed via %s (%s)",
		route.Region,
		route.Node,
	)

	msg := new(dns.Msg)
	msg.SetReply(req)

	if len(req.Question) > 0 {

		name := strings.TrimSuffix(
			req.Question[0].Name,
			".",
		)

		record := repository.Find(name)

		if record.Type == "A" {

			ip := ResolveSteeredIP()

			AddARecord(
				msg,
				req.Question[0].Name,
				ip,
				record.TTL,
			)
		}
	}

	_ = w.WriteMsg(msg)
}
