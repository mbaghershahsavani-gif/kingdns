package resolver

import (
	"strings"

	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/controller"
	"github.com/miekg/dns"
)

var authenticatedController = controller.NewClient("http://localhost:8080")

func AuthControllerHandler(w dns.ResponseWriter, req *dns.Msg) {
	msg := new(dns.Msg)
	msg.SetReply(req)

	if len(req.Question) > 0 {
		name := strings.TrimSuffix(req.Question[0].Name, ".")

		if name == "kingdns.local" {
			nodes, _ := authenticatedController.GetAuthenticatedNodes()

			if len(nodes) > 0 {
				AddARecord(msg, req.Question[0].Name, "127.0.0.1", 60)
			}
		}
	}

	_ = w.WriteMsg(msg)
}
