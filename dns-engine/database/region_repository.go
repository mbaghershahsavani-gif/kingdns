package database

type RegionRecord struct {
	Name string
	Node string
}

func GetRegions() []RegionRecord {
	return []RegionRecord{
		{
			Name: "local",
			Node: "127.0.0.1",
		},
	}
}
