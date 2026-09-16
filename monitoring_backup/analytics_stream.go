package monitoring

func Record(event AnalyticsEvent) bool {
	return event.Name != ""
}
