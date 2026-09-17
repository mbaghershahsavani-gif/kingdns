package cluster

import (
	"os"
	"time"

	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/version"
)

type Node struct {
	ID           string    `json:"node_id"`
	Region       string    `json:"region"`
	IP           string    `json:"ip"`
	Version      string    `json:"version"`
	Status       string    `json:"status"`
	Capabilities []string  `json:"capabilities"`
	LastSeen     time.Time `json:"last_seen"`
}

func CurrentNode() Node {

	region := os.Getenv("KINGDNS_REGION")

	if region == "" {
		region = "unknown"
	}

	nodeID := os.Getenv("KINGDNS_NODE_ID")
	ip := os.Getenv("KINGDNS_NODE_IP")

	if ip == "" {
		ip = "127.0.0.1"
	}

	if nodeID == "" {
		nodeID = region + "-01"
	}

	return Node{
		ID:      nodeID,
		Region:  region,
		IP:      ip,
		Version: version.Version,
		Status:  "healthy",
		Capabilities: []string{
			"dns",
			"udp",
			"tcp",
		},
		LastSeen: time.Now(),
	}
}
