package resolver

import "github.com/mbaghershahsavani-gif/kingdns-dns-engine/routing"

func ResolveIntelligent(domain string) string {
	nodes := []routing.IntelligentCandidate{
		{
			IP:    "127.0.0.1",
			Score: 100,
		},
	}

	return routing.SelectOptimal(nodes)
}
