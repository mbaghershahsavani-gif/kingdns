package resolver

import "github.com/mbaghershahsavani-gif/kingdns-dns-engine/internal/cluster"

func ResolveSteeredIP() string {

	route := cluster.CurrentRoute()

	switch route.Region {

	case "iran":
		return "10.10.10.10"

	case "international":
		return "20.20.20.20"

	default:
		return "20.20.20.20"
	}
}
