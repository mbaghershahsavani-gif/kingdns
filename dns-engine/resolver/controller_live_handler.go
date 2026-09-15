package resolver

import (
	"strings"

	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/controller"
	"github.com/miekg/dns"
)

var controllerClient = controller.NewClient("http://localhost:8080")

func ControllerLiveHandler(w dns.ResponseWriter, req *dns.Msg) {
	msg := new(dns.Msg)
	msg.SetReply(req)

	if len(req.Question) > 0 {
		_ = strings.TrimSuffix(req.Question[0].Name, ".")

		nodes, _ := controllerClient.GetNodes()

		if len(nodes) > 0 {
			AddARecord(
				msg,
				req.Question[0].Name,
				"127.0.0.1",
				60,
			)
		}
	}

	_ = w.WriteMsg(msg)
}
