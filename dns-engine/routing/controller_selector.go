package routing

import (
	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/controller"
)

func SelectControllerNode(nodes []controller.Node) controller.Node {
	for _, node := range nodes {
		if node.Status == "online" {
			return node
		}
	}

	return controller.Node{}
}
