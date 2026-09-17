package cluster

import (
	"log"
	"time"
)

func StartMonitor() {

	ticker := time.NewTicker(
		15 * time.Second,
	)

	defer ticker.Stop()

	for range ticker.C {

		UpdateHeartbeat()

		RegisterScore(
			CalculateScore(),
		)

		SyncPeers()

		RemoveStaleNodes()

		hb := GetHeartbeat()

		log.Printf(
			"Heartbeat sent: %s %s %s",
			hb.NodeID,
			hb.Region,
			hb.Status,
		)
	}
}
