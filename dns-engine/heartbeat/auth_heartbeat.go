package heartbeat

type AuthHeartbeat struct {
	Token  string
	Node   string
	Status string
}

func Send(h AuthHeartbeat) {
}
