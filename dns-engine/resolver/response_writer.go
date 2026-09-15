package resolver

import "github.com/miekg/dns"

func WriteResponse(w dns.ResponseWriter, request *dns.Msg, record DNSRecord) error {
	response := new(dns.Msg)
	response.SetReply(request)

	// Future:
	// - build DNS A/AAAA/CNAME answers
	// - write packet response

	return w.WriteMsg(response)
}
