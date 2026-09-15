package database

type DomainRepository struct {
	DB Client
}

func (r DomainRepository) Find(name string) DomainRecord {

	if name == "kingdns.local" {
		return DomainRecord{
			Name:  name,
			Type:  "A",
			Value: "127.0.0.1",
			TTL:   60,
		}
	}

	return DomainRecord{}
}