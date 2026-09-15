package monitoring

type QueryMetrics struct {
	Total     uint64
	CacheHits uint64
}

func (m *QueryMetrics) RecordQuery() {
	m.Total++
}
