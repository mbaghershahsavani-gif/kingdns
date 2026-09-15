package resolver

import (
	"github.com/miekg/dns"
)

func BuildResponse(req *dns.Msg, domain string, ip string) *dns.Msg {
	msg := new(dns.Msg)
	msg.SetReply(req)

	AddARecord(msg, domain, ip, 60)

	return msg
}
