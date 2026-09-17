package cluster

import (
	"log"
	"time"
)

func RemoveStaleNodes() {

	nodesMutex.Lock()
	defer nodesMutex.Unlock()

	now := time.Now()

	for id, node := range nodes {

		age := now.Sub(node.LastSeen)

		if age > nodeTimeout {

			log.Printf(
				"Removing stale node: %s age=%s",
				id,
				age,
			)

			delete(nodes, id)
		}
	}
}
