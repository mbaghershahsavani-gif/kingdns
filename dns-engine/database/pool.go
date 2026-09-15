package database

type Pool struct {
	Connected bool
}

func NewPool() *Pool {
	return &Pool{Connected: true}
}

func (p *Pool) Healthy() bool {
	return p.Connected
}
