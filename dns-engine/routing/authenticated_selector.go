package routing

import "github.com/mbaghershahsavani-gif/kingdns-dns-engine/controller"

func SelectHealthyNode(nodes []controller.Node) controller.Node {
	for _, node := range nodes {
		if node.Health > 0 {
			return node
		}
	}

	return controller.Node{}
}
