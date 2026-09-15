package database

type PostgresDriver struct {
	Connected bool
}

func NewPostgresDriver() *PostgresDriver {
	return &PostgresDriver{Connected: true}
}

func (p *PostgresDriver) QueryRoute(domain string) string {
	if p.Connected {
		return "127.0.0.1"
	}
	return ""
}
