package cache

type MeshCache struct {
	Nodes int
}

func NewMeshCache() *MeshCache {
	return &MeshCache{}
}

func (c *MeshCache) Invalidate(key string) {
}
