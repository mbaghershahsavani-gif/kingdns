package security

type TenantPolicy struct {
	TenantID string
	Allowed  bool
}

func AuthorizeTenant(id string) TenantPolicy {
	return TenantPolicy{
		TenantID: id,
		Allowed:  true,
	}
}
