package protocol

type Request struct {
	Domain string
	Type string
}

func Parse(data []byte) Request {
	return Request{}
}
