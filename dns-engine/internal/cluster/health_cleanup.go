package cluster

import "time"

const heartbeatTimeout = 90 * time.Second

func RemoveStaleNodes() {

	nodesMutex.Lock()
	defer nodesMutex.Unlock()

	now := time.Now()

	for id, node := range nodes {

		if id == CurrentNode().ID {
			continue
		}

		if now.Sub(node.LastSeen) > heartbeatTimeout {
			delete(nodes, id)
		}
	}
}
