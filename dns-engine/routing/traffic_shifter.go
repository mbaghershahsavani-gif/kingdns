package routing

type TrafficShift struct {
	Region string
	Weight int
}

func ShiftTraffic(region string, weight int) TrafficShift {
	return TrafficShift{
		Region: region,
		Weight: weight,
	}
}
