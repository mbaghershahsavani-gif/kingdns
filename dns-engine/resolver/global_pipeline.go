package resolver

import "github.com/mbaghershahsavani-gif/kingdns-dns-engine/routing"

func ResolveGlobal(domain string) string {
	nodes := []routing.FailoverNode{
		{
			IP:      "127.0.0.1",
			Healthy: true,
		},
	}

	return routing.SelectFailover(nodes)
}
