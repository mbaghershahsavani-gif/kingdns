package cache

type GlobalStateCache struct {
	Synchronized bool
}

func NewGlobalStateCache() *GlobalStateCache {
	return &GlobalStateCache{
		Synchronized: true,
	}
}
