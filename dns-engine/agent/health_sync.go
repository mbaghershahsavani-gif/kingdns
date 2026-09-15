package agent

type Health struct {
	NodeID int
	Score  int
}

func Update(health Health) {
	// Future:
	// consume agent health information
}
