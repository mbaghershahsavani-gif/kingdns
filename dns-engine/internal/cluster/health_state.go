package cluster

import "time"

type HealthState string

const (
	Healthy  HealthState = "healthy"
	Degraded HealthState = "degraded"
	Offline  HealthState = "offline"
)

func EvaluateHealth(score Score) HealthState {

	if score.Heartbeat == 0 {
		return Offline
	}

	if score.Total < 50 {
		return Degraded
	}

	hb := GetHeartbeat()

	if time.Since(hb.LastSeen) > time.Minute {
		return Offline
	}

	return Healthy
}
