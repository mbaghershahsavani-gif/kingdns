package cache

type TTLManager struct {
}

func (t TTLManager) Expired(ttl int) bool {
	return ttl <= 0
}
