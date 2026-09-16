package database

import (
	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/internal/cluster"
)

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

	if name == "example.com" {

		route := cluster.CurrentRoute()

		switch route.Region {

		case "iran":

			return DomainRecord{
				Name:  name,
				Type:  "A",
				Value: "10.10.10.10",
				TTL:   60,
			}

		case "international":

			return DomainRecord{
				Name:  name,
				Type:  "A",
				Value: "20.20.20.20",
				TTL:   60,
			}
		}
	}

	return DomainRecord{}
}
