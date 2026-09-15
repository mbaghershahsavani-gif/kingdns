package monitoring

type AnalyticsEvent struct {
	Name  string
	Value int
}

func Record(event AnalyticsEvent) bool {
	return event.Name != ""
}
