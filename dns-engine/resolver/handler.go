package resolver

func Handle(domain string) Response {
	return Response{
		Domain: domain,
		TTL: 60,
	}
}
