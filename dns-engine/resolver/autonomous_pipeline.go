package resolver

import "github.com/mbaghershahsavani-gif/kingdns-dns-engine/routing"

func ResolveAutonomous(domain string) string {
	nodes := []routing.AutonomousNode{
		{
			IP:    "127.0.0.1",
			Score: 100,
		},
	}

	return routing.SelectAutonomous(nodes)
}
