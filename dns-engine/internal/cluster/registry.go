package cluster

import (
	"log"
	"sync"
)

var (
	nodes      = make(map[string]Score)
	nodesMutex sync.RWMutex
)

func RegisterScore(score Score) {

	if score.NodeID == "" {
		return
	}

	if !IsTrusted(score.NodeID) {
		log.Printf("Rejected untrusted node: %s", score.NodeID)
		return
	}

	log.Printf("Registered trusted node: %s", score.NodeID)

	nodesMutex.Lock()
	defer nodesMutex.Unlock()

	nodes[score.NodeID] = score
}

func RemoveNode(id string) {

	nodesMutex.Lock()
	defer nodesMutex.Unlock()

	delete(nodes, id)
}

func GetNodes() []Score {

	nodesMutex.RLock()
	defer nodesMutex.RUnlock()

	result := make([]Score, 0, len(nodes))

	for _, node := range nodes {
		result = append(result, node)
	}

	return result
}

func ClearNodes() {

	nodesMutex.Lock()
	defer nodesMutex.Unlock()

	nodes = make(map[string]Score)
}
