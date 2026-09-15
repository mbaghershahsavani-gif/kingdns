package resolver

import (
	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/cache"
	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/database"
)

func ResolveControlPlane(domain string) string {
	cache := cache.NewDistributedCache()

	if value := cache.Get(domain); value != "" {
		return value
	}

	db := database.NewPGXPool()
	target := db.Query(domain)

	if target != "" {
		cache.Set(domain, target, 60)
	}

	return target
}
