package cluster

import "time"

type Score struct {
	NodeID string `json:"node_id"`
	Region string `json:"region"`

	Status string `json:"status"`

	Heartbeat int `json:"heartbeat"`
	DNS       int `json:"dns"`
	Latency   int `json:"latency"`

	Total int `json:"total"`
}

func CalculateScore() Score {

	hb := GetHeartbeat()

	heartbeatScore := 0

	if time.Since(hb.LastSeen) < time.Minute {
		heartbeatScore = 40
	}

	dnsScore := 30
	latencyScore := 10

	return Score{
		NodeID: hb.NodeID,
		Region: hb.Region,
		Status: hb.Status,

		Heartbeat: heartbeatScore,
		DNS:       dnsScore,
		Latency:   latencyScore,

		Total: heartbeatScore +
			dnsScore +
			latencyScore,
	}
}
