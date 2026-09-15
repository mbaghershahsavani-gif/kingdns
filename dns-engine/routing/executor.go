package routing

type Result struct {
	Target string
	TTL int
}

func Execute() Result {
	return Result{
		TTL: 60,
	}
}
