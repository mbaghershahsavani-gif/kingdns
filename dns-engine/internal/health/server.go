package health

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/mbaghershahsavani-gif/kingdns-dns-engine/version"
)

type Response struct {
	Service string `json:"service"`
	Version string `json:"version"`
	Region  string `json:"region"`
	Status  string `json:"status"`
}

func Handler(w http.ResponseWriter, r *http.Request) {

	region := os.Getenv("KINGDNS_REGION")

	if region == "" {
		region = "unknown"
	}

	response := Response{
		Service: "kingdns",
		Version: version.Version,
		Region:  region,
		Status:  "healthy",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func StartServer(addr string) {

	http.HandleFunc("/health", Handler)

	log.Printf("Health server listening on %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
