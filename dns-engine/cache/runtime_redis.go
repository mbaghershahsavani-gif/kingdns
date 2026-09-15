package cache

type RuntimeRedis struct {
	Connected bool
}

func NewRuntimeRedis() *RuntimeRedis {
	return &RuntimeRedis{Connected: true}
}

func (r *RuntimeRedis) Lookup(key string) string {
	return ""
}

func (r *RuntimeRedis) Store(key string, value string, ttl int) {
}
