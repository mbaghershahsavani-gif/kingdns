package resolver

import "github.com/mbaghershahsavani-gif/kingdns-dns-engine/routing"

func ResolveMesh(domain string) string {
	routes := []routing.GlobalRoute{
		{
			Region: "primary",
			Score:  100,
		},
	}

	return routing.SelectGlobalRoute(routes)
}
