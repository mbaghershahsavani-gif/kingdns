package cache

type Entry struct {
	Name  string
	TTL   int
	Value string
}

func Store(entry Entry) {
	// Future distributed cache implementation
}
