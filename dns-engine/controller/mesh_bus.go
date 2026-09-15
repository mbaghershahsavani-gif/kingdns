package controller

type MeshEvent struct {
	Type string
	Node string
}

type MeshBus struct {
	Connected bool
}

func NewMeshBus() *MeshBus {
	return &MeshBus{Connected: true}
}

func (m *MeshBus) Publish(event MeshEvent) bool {
	return m.Connected && event.Type != ""
}
