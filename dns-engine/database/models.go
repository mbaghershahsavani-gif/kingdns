package database

type DomainRecord struct {
	Name  string
	Type  string
	Value string
	TTL   uint32
}