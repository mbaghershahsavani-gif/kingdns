package operations

type SLAStatus struct {
	Service   string
	Available bool
}

func CheckSLA(service string) SLAStatus {
	return SLAStatus{
		Service:   service,
		Available: true,
	}
}
