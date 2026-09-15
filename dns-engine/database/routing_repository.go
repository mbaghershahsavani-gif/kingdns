package database

type RoutingRule struct {
	Domain   string
	NodeID   int
	Priority int
}

func GetRoutingRules() []RoutingRule {
	return []RoutingRule{}
}
