package cluster

func BootstrapTrust() {

	AddTrustedNode(
		TrustedNode{
			NodeID:  "int-01",
			Region:  "international",
			Trusted: true,
		},
	)

	AddTrustedNode(
		TrustedNode{
			NodeID:  "ir-01",
			Region:  "iran",
			Trusted: true,
		},
	)
}
