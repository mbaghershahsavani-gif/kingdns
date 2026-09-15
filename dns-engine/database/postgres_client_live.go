package database

type PostgreSQLClient struct {
	Connected bool
}

func NewPostgreSQLClient() *PostgreSQLClient {
	return &PostgreSQLClient{Connected: true}
}

func (c *PostgreSQLClient) FindTarget(domain string) string {
	if c.Connected {
		return "127.0.0.1"
	}

	return ""
}
