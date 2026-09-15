package cache

var records = map[string]Record{}

func Get(name string) (Record, bool) {
	record, ok := records[name]
	return record, ok
}

func Set(record Record) {
	records[record.Name] = record
}
