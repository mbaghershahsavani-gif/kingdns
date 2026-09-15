package resolver

import "github.com/mbaghershahsavani-gif/kingdns-dns-engine/routing"

func ResolveKubernetes(domain string) string {
	nodes := []routing.ServiceEndpoint{
		{
			IP:      "127.0.0.1",
			Healthy: true,
		},
	}

	return routing.SelectServiceEndpoint(nodes)
}
