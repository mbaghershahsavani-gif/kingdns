package validation

import "errors"

type NodeInput struct {
	Name string
	Type string
	Country string
	IPAddress string
}

func ValidateNode(n NodeInput) error {
	if n.Name == "" {
		return errors.New("node name is required")
	}
	if n.Type == "" {
		return errors.New("node type is required")
	}
	if n.Country == "" {
		return errors.New("node country is required")
	}
	if n.IPAddress == "" {
		return errors.New("node ip address is required")
	}
	return nil
}
