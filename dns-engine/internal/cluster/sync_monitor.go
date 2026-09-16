package cluster

import (
	"log"
	"time"
)

func StartSyncMonitor() {

	ticker := time.NewTicker(
		60 * time.Second,
	)

	defer ticker.Stop()

	for {

		select {

		case <-ticker.C:

			SyncPeers()

			log.Println(
				"Cluster peer synchronization completed",
			)
		}
	}
}
