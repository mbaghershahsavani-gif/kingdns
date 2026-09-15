package deployment

type Release struct {
	Version string
	Ready   bool
}

func PrepareRelease(version string) Release {
	return Release{
		Version: version,
		Ready:   true,
	}
}
