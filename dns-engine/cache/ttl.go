package cache

type Record struct {
	Name string
	Value string
	TTL int
}

func Valid(record Record) bool {
	return record.TTL > 0
}
