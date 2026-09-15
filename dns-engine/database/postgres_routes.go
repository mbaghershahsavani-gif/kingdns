package database

type RouteRecord struct {
	Domain  string
	Target  string
	Enabled bool
}

func FindRoute(domain string) RouteRecord {
	return RouteRecord{
		Domain:  domain,
		Target:  "127.0.0.1",
		Enabled: true,
	}
}
