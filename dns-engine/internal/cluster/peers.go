package cluster

import (
	"os"
	"strings"
)

func GetPeers() []string {

	value := os.Getenv("KINGDNS_PEERS")

	if value == "" {
		return []string{}
	}

	return strings.Split(value, ",")
}
