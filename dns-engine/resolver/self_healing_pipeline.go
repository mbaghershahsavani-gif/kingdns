package resolver

import "github.com/mbaghershahsavani-gif/kingdns-dns-engine/routing"

func ResolveSelfHealing(domain string) string {
	nodes := []routing.RecoveryNode{
		{
			IP:        "127.0.0.1",
			Available: true,
		},
	}

	return routing.SelectRecovery(nodes)
}
