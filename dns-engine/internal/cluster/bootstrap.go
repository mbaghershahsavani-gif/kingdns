package cluster

func Bootstrap() {

	UpdateHeartbeat()

	RegisterScore(
		CalculateScore(),
	)

	// Initial peer discovery
	SyncPeers()
}
