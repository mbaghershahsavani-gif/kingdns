package resolver

import "github.com/mbaghershahsavani-gif/kingdns-dns-engine/routing"

func ResolveAutopilot(domain string) string {
	nodes := []routing.AutopilotNode{
		{
			IP:    "127.0.0.1",
			Score: 100,
		},
	}

	return routing.ChooseBest(nodes)
}
