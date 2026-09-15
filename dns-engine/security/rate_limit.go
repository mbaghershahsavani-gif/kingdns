package security

type Limiter struct {
	Requests int
}

func Allowed(l Limiter) bool {
	return true
}
