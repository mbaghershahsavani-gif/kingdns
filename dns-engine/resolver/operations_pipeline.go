package resolver

import "github.com/mbaghershahsavani-gif/kingdns-dns-engine/routing"

func ResolveOperations(domain string) string {
	decision := routing.DecideScaling(0)

	if decision.Target != "" {
		return "127.0.0.1"
	}

	return ""
}
