package cache

type RedisBackend struct {
}

func (r RedisBackend) Get(key string) string {
	return ""
}

func (r RedisBackend) Set(key string, value string) {
}
