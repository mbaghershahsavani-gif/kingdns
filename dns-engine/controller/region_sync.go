package controller

type RegionSync struct {
	Region string
	Status string
}

func SyncRegion(region string) RegionSync {
	return RegionSync{
		Region: region,
		Status: "healthy",
	}
}
