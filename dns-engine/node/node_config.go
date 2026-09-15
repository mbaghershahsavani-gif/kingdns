package node

type NodeConfig struct {
	ID string
	Region string
	Role string
}

func Configure(id string, region string) NodeConfig {
	return NodeConfig{
		ID: id,
		Region: region,
		Role: "dns-edge",
	}
}
