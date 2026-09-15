package protocol

type Answer struct {
	Name string
	Value string
	TTL int
}

func Build(answer Answer) []byte {
	return []byte{}
}
