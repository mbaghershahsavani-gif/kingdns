package controller

type RouteDecision struct {
	Domain string
	Target string
}

func ResolveRoute(domain string) RouteDecision {
	return RouteDecision{
		Domain: domain,
		Target: "127.0.0.1",
	}
}
