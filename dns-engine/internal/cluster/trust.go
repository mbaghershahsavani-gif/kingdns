package cluster

import "sync"

type TrustedNode struct {
	NodeID  string `json:"node_id"`
	Region  string `json:"region"`
	Trusted bool   `json:"trusted"`
}

var (
	trustedNodes = map[string]TrustedNode{}
	trustMutex   sync.RWMutex
)

func AddTrustedNode(node TrustedNode) {

	trustMutex.Lock()
	defer trustMutex.Unlock()

	trustedNodes[node.NodeID] = node
}

func IsTrusted(nodeID string) bool {

	trustMutex.RLock()
	defer trustMutex.RUnlock()

	node, exists := trustedNodes[nodeID]

	return exists && node.Trusted
}

func GetTrustedNodes() []TrustedNode {

	trustMutex.RLock()
	defer trustMutex.RUnlock()

	result := make([]TrustedNode, 0, len(trustedNodes))

	for _, node := range trustedNodes {
		result = append(result, node)
	}

	return result
}
