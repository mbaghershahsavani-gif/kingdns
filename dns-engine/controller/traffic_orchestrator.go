package controller

type TrafficMigration struct {
	From string
	To   string
}

func MigrateTraffic(from string, to string) TrafficMigration {
	return TrafficMigration{
		From: from,
		To:   to,
	}
}
