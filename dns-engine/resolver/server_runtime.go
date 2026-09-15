package resolver

import "github.com/miekg/dns"

func StartRuntime() {
	dns.HandleFunc(".", RuntimeHandler)
}
