package heartbeat

type Report struct {
	NodeID uint `json:"node_id"`
	Latency int `json:"latency"`
	CPU int `json:"cpu"`
	Memory int `json:"memory"`
	Status string `json:"status"`
}
