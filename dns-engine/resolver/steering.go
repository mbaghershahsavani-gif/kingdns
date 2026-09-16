package resolver

import (
	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/internal/cluster"
	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/policies"
)

func ResolveSteeredIP(domain string) string {

	route := cluster.CurrentRoute()

	policy, ok := policies.Find(domain)

	if !ok {
		return ""
	}

	switch route.Region {

	case "iran":
		return policy.IranIP

	case "international":
		return policy.InternationalIP

	default:
		return policy.InternationalIP
	}
}
