package cluster

import (
	"log"
	"time"
)

func StartMonitor() {

	ticker := time.NewTicker(
		30 * time.Second,
	)

	defer ticker.Stop()

	for {
		select {

		case <-ticker.C:

			UpdateHeartbeat()

			RegisterScore(
				CalculateScore(),
			)

			hb := GetHeartbeat()

			log.Printf(
				"Heartbeat sent: %s %s %s",
				hb.NodeID,
				hb.Region,
				hb.Status,
			)

		}
	}
}
