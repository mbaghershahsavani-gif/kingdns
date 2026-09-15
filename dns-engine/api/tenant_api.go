package api

type TenantRequest struct {
	Name string
}

type TenantResponse struct {
	ID string
}

func CreateTenant(req TenantRequest) TenantResponse {
	return TenantResponse{
		ID: "tenant-id",
	}
}
