package security

type ComplianceReport struct {
	Name   string
	Passed bool
}

func GenerateReport(name string) ComplianceReport {
	return ComplianceReport{
		Name:   name,
		Passed: true,
	}
}
