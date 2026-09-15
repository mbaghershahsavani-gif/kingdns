package intelligence

type Anomaly struct {
	Metric   string
	Detected bool
}

func Detect(metric string) Anomaly {
	return Anomaly{
		Metric:   metric,
		Detected: false,
	}
}
