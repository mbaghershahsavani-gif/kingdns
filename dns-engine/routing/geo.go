package routing

type Location struct {
	Country string
	Region string
}

func Match(location Location) Candidate {
	return Candidate{}
}
