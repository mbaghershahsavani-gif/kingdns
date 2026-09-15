package controller

type RouteResult struct {
	Domain string
	Target string
}

func ExecuteRoute(domain string) RouteResult {
	return RouteResult{
		Domain: domain,
		Target: "127.0.0.1",
	}
}
