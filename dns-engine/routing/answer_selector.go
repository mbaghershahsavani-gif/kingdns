package routing

type Selection struct {
	Target string
}

func SelectAnswer(domain string) Selection {
	return Selection{
		Target: "127.0.0.1",
	}
}
