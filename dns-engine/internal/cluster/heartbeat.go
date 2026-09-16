package cluster

import (
	"sync"
	"time"
)

type Heartbeat struct {
	NodeID   string    `json:"node_id"`
	Region   string    `json:"region"`
	Status   string    `json:"status"`
	LastSeen time.Time `json:"last_seen"`
}

var (
	CurrentHeartbeat Heartbeat
	Mutex            sync.RWMutex
)

func UpdateHeartbeat() {

	node := CurrentNode()

	Mutex.Lock()
	defer Mutex.Unlock()

	CurrentHeartbeat = Heartbeat{
		NodeID:   node.ID,
		Region:   node.Region,
		Status:   "healthy",
		LastSeen: time.Now(),
	}
}

func GetHeartbeat() Heartbeat {

	Mutex.RLock()
	defer Mutex.RUnlock()

	if CurrentHeartbeat.NodeID == "" {

		Mutex.RUnlock()

		UpdateHeartbeat()

		Mutex.RLock()
	}

	return CurrentHeartbeat
}
