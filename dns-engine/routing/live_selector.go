package routing

import "github.com/mbaghershahsavani-gif/kingdns-dns-engine/controller"

func SelectLiveNode(nodes []controller.APINode) controller.APINode {
	for _, node := range nodes {
		if node.Status == "online" && node.Health > 0 {
			return node
		}
	}

	return controller.APINode{}
}
