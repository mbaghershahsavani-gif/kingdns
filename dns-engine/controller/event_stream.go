package controller

type Event struct {
	Type string
	Node string
}

func Publish(event Event) bool {
	return event.Type != ""
}
