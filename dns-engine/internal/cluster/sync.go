package cluster

import (
	"encoding/json"
	"log"
	"net/http"
)

func SyncHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	nodes := GetNodes()

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(nodes)

	log.Printf(
		"Cluster sync request served: nodes=%d",
		len(nodes),
	)
}
