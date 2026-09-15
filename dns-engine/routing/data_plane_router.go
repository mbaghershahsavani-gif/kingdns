package routing

import (
	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/database"
)

func Resolve(domain string) string {
	route := database.FindRoute(domain)

	if route.Enabled {
		return route.Target
	}

	return ""
}
