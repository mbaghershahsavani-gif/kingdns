package controlplane

type FederationNode struct {
	Region  string
	Healthy bool
}

func RegisterRegion(region string) FederationNode {
	return FederationNode{
		Region:  region,
		Healthy: true,
	}
}
