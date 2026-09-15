package database

type TenantRecord struct {
	ID   string
	Name string
}

func SaveTenant(record TenantRecord) TenantRecord {
	return record
}
