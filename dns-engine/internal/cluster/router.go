package cluster

type Route struct {
	Region string `json:"region"`
	Node   string `json:"node"`
	State  string `json:"state"`
}

func CurrentRoute() Route {

	decision := CalculateDecision()

	return Route{
		Region: decision.ActiveRegion,
		Node:   decision.ActiveNode,
		State:  decision.State,
	}
}
