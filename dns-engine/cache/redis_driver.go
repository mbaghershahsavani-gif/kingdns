package cache

type RedisDriver struct {
	Connected bool
}

func NewRedisDriver() *RedisDriver {
	return &RedisDriver{Connected: true}
}

func (r *RedisDriver) Get(key string) string {
	return ""
}

func (r *RedisDriver) Set(key string, value string, ttl int) {
}
