package monitoring

type RecoveryEvent struct {
	Node  string
	Event string
}

func RecordRecovery(node string) RecoveryEvent {
	return RecoveryEvent{
		Node:  node,
		Event: "recovery-started",
	}
}
