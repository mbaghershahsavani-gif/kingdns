package database

type DomainQuery struct {
	Name  string
	Type  string
	Value string
	TTL   uint32
}

func (q DomainQuery) Find(name string) DomainQuery {
	// v3.0 foundation.
	// Next step: execute PostgreSQL query.

	if name == "kingdns.local" {
		return DomainQuery{
			Name:  name,
			Type:  "A",
			Value: "127.0.0.1",
			TTL:   60,
		}
	}

	return DomainQuery{}
}
