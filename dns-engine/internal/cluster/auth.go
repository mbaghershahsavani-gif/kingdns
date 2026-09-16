package cluster

import (
	"net/http"
	"os"
)

func Authenticate(
	r *http.Request,
) bool {

	expected :=
		os.Getenv("KINGDNS_CLUSTER_TOKEN")

	if expected == "" {
		return true
	}

	token :=
		r.Header.Get("Authorization")

	return token ==
		"Bearer "+expected
}
