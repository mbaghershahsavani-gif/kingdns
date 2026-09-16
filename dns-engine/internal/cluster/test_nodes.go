package cluster

import (
	"encoding/json"
	"net/http"
)

func AddTestNodeHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	testNode := Score{
		NodeID: "ir-01",
		Region: "iran",
		Status: "healthy",

		Heartbeat: 40,
		DNS:       30,
		Latency:   10,

		Total: 95,
	}

	RegisterScore(testNode)

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(testNode)
}

func RemoveTestNodeHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	RemoveNode("ir-01")

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]string{
			"removed": "ir-01",
		},
	)
}
