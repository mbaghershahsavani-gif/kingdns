package protocol

type RecordType uint16

const (
	A RecordType = 1
	AAAA RecordType = 28
	CNAME RecordType = 5
)

type Question struct {
	Name string
	Type RecordType
}
