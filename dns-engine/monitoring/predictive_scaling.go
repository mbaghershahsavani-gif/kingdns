package monitoring

type ScalingForecast struct {
	ExpectedQueries int
	ScaleRequired   bool
}

func ForecastScaling(queries int) ScalingForecast {
	return ScalingForecast{
		ExpectedQueries: queries,
		ScaleRequired:   queries > 10000,
	}
}
