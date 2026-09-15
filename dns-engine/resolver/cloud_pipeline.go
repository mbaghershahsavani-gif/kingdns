package resolver

import "github.com/mbaghershahsavani-gif/kingdns-dns-engine/routing"

func ResolveCloud(domain string) string {
	nodes := []routing.InfrastructureNode{
		{
			IP:       "127.0.0.1",
			Capacity: 100,
			Healthy:  true,
		},
	}

	return routing.SelectInfrastructureNode(nodes)
}
