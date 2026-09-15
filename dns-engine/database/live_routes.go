package database

type LiveRoute struct {
	Domain  string
	Target  string
	Enabled bool
}

func FindLiveRoute(domain string) LiveRoute {
	return LiveRoute{
		Domain:  domain,
		Target:  "127.0.0.1",
		Enabled: true,
	}
}
