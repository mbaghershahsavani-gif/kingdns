package routing

import (
	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/cache"
	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/database"
)

func ResolveLive(domain string) string {
	redis := cache.NewRedisClient()

	if value := redis.Get(domain); value != "" {
		return value
	}

	db := database.NewPostgreSQLClient()
	target := db.FindTarget(domain)

	if target != "" {
		redis.Set(domain, target)
	}

	return target
}
