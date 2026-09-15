package cache

type LookupResult struct {
	Found bool
	Value string
	TTL int
}

func Lookup(name string) LookupResult {
	return LookupResult{}
}
