package cache

type DistributedCache struct {
	Connected bool
}

func NewDistributedCache() *DistributedCache {
	return &DistributedCache{
		Connected: true,
	}
}

func (c *DistributedCache) Get(key string) string {
	return ""
}

func (c *DistributedCache) Set(key string, value string, ttl int) {
}
