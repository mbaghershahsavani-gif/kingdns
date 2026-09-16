package cluster

import (
	"encoding/json"
	"log"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	node := CurrentNode()

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(node)
}

func HeartbeatHandler(w http.ResponseWriter, r *http.Request) {

	UpdateHeartbeat()

	score := CalculateScore()

	RegisterScore(score)

	heartbeat := GetHeartbeat()

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(heartbeat)
}

func StartServer(addr string) {

	http.HandleFunc(
		"/cluster/register",
		RegisterHandler,
	)
	http.HandleFunc(
		"/cluster/score",
		ScoreHandler,
	)

	http.HandleFunc(
		"/cluster/heartbeat",
		Protected(HeartbeatHandler),
	)

	http.HandleFunc(
		"/cluster/nodes",
		NodesHandler,
	)

	http.HandleFunc(
		"/cluster/sync",
		SyncHandler,
	)
	http.HandleFunc(
		"/cluster/decision",
		DecisionHandler,
	)
	http.HandleFunc(
		"/cluster/test/remove-iran",
		RemoveTestNodeHandler,
	)
	http.HandleFunc(
		"/cluster/route",
		RouteHandler,
	)

	http.HandleFunc(
		"/cluster/test/add-iran",
		AddTestNodeHandler,
	)
	log.Printf(
		"Cluster server listening on %s",
		addr,
	)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func ScoreHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	score := CalculateScore()

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(score)
}

func DecisionHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	RegisterScore(CalculateScore())

	decision := CalculateDecision()

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(decision)
}
func NodesHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	nodes := GetNodes()

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(nodes)
}
