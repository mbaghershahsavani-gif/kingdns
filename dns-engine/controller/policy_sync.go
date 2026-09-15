package controller

type PolicySync struct {
	Status string
}

func SyncPolicies() PolicySync {
	return PolicySync{
		Status: "updated",
	}
}
