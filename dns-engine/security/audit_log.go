package security

type AuditEvent struct {
	User   string
	Action string
}

func RecordAudit(user string, action string) AuditEvent {
	return AuditEvent{
		User:   user,
		Action: action,
	}
}
