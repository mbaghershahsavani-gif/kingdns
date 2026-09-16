package cluster

import (
	"encoding/json"
	"net/http"
)

func RouteHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	route := CurrentRoute()

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(route)
}
