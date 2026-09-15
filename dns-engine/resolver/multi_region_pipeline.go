package resolver

import (
	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/routing"
)

func ResolveMultiRegion(domain string) string {
	nodes := []routing.RegionNode{
		{
			Region:  "local",
			IP:      "127.0.0.1",
			Latency: 1,
			Healthy: true,
		},
	}

	return routing.SelectRegion(nodes)
}
