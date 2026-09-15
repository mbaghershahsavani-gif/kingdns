package tenant

type Zone struct {
	TenantID string
	Domain   string
}

func CreateZone(tenantID string, domain string) Zone {
	return Zone{
		TenantID: tenantID,
		Domain:   domain,
	}
}
