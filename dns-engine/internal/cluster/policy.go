package cluster

import "os"

func PreferredRegion() string {

	region := os.Getenv("KINGDNS_PREFERRED_REGION")

	if region == "" {
		return "iran"
	}

	return region
}
