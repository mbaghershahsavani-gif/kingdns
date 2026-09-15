package resolver

import (
	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/cache"
	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/database"
)

func ResolveRuntime(domain string) string {
	redis := cache.NewRuntimeRedis()

	if value := redis.Lookup(domain); value != "" {
		return value
	}

	db := database.NewRuntimePostgres()
	target := db.LookupRoute(domain)

	if target != "" {
		redis.Store(domain, target, 60)
	}

	return target
}
