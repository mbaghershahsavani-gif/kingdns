package database

type PGXPool struct {
	Connected bool
}

func NewPGXPool() *PGXPool {
	return &PGXPool{Connected: true}
}

func (p *PGXPool) Query(domain string) string {
	if p.Connected {
		return "127.0.0.1"
	}
	return ""
}
