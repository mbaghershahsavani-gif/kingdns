package metrics

type SystemMetrics struct {
    CPU float64
    Memory float64
    Uptime int64
}

func Collect() SystemMetrics {
    return SystemMetrics{}
}
