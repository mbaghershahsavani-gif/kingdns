package resolver

import "github.com/mbaghershahsavani-gif/kingdns-dns-engine/routing"

func ResolveMesh(domain string) string {
	nodes := []routing.MeshNode{
		{
			Region: "local",
			IP:     "127.0.0.1",
			Score:  100,
		},
	}

	return routing.SelectMeshNode(nodes)
}
