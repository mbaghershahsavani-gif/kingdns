package routing

type RuntimeSelector struct {
}

func (s RuntimeSelector) Select() Result {
	return Result{
		TTL: 60,
	}
}
