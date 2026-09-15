package operations

type RecoveryPlan struct {
	Name  string
	Ready bool
}

func CreateRecoveryPlan(name string) RecoveryPlan {
	return RecoveryPlan{
		Name:  name,
		Ready: true,
	}
}
