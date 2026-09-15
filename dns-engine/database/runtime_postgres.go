package database

type RuntimePostgres struct {
	Connected bool
}

func NewRuntimePostgres() *RuntimePostgres {
	return &RuntimePostgres{Connected: true}
}

func (p *RuntimePostgres) LookupRoute(domain string) string {
	if p.Connected {
		return "127.0.0.1"
	}
	return ""
}
