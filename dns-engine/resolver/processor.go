package resolver

type Processor struct {
}

func (p Processor) Resolve(query Query) Response {
	return Response{
		Domain: query.Domain,
		TTL: 60,
	}
}
