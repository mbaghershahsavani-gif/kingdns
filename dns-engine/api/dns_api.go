package api

type DNSRequest struct {
	Domain string
}

type DNSResponse struct {
	Address string
}

func Resolve(request DNSRequest) DNSResponse {
	return DNSResponse{
		Address: "127.0.0.1",
	}
}
