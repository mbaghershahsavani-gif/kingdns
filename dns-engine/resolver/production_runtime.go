package resolver

import (
	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/cache"
	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/database"
)

func ResolveProduction(domain string) string {
	redis := cache.NewRedisDriver()

	if value := redis.Get(domain); value != "" {
		return value
	}

	db := database.NewPostgresDriver()
	return db.QueryRoute(domain)
}
