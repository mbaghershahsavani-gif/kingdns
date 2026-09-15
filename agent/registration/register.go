package registration

import (
    "github.com/mbaghershahsavani-gif/kingdns-agent/client"
)

type NodeRegistration struct {
    Name string `json:"name"`
    Type string `json:"type"`
    Region string `json:"region"`
}

func Register(controller string, node NodeRegistration) error {
    return client.PostJSON(
        controller+"/api/nodes/register",
        node,
    )
}
