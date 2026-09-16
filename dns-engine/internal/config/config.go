package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	NodeID string `json:"node_id"`
	Region string `json:"region"`

	ClusterToken string `json:"cluster_token"`

	PreferredRegion string `json:"preferred_region"`

	ClusterPort int `json:"cluster_port"`
	HealthPort  int `json:"health_port"`
	DNSPort     int `json:"dns_port"`
}

func Load(path string) (Config, error) {

	data, err := os.ReadFile(path)

	if err != nil {
		return Config{}, err
	}

	var cfg Config

	err = json.Unmarshal(
		data,
		&cfg,
	)

	return cfg, err
}
