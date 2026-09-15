package operations

type Incident struct {
	ID     string
	Status string
}

func CreateIncident(id string) Incident {
	return Incident{
		ID:     id,
		Status: "open",
	}
}
