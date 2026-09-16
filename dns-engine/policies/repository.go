package policies

import (
	"encoding/json"
	"os"
	"sync"
)

type Policy struct {
	Domain          string `json:"domain"`
	Type            string `json:"type"`
	IranIP          string `json:"iran_ip"`
	InternationalIP string `json:"international_ip"`
	TTL             uint32 `json:"ttl"`
}

var (
	policies = make(map[string]Policy)
	mutex    sync.RWMutex
)

func Load(path string) error {

	data, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	var policy Policy

	err = json.Unmarshal(data, &policy)

	if err != nil {
		return err
	}

	mutex.Lock()
	defer mutex.Unlock()

	policies[policy.Domain] = policy

	return nil
}

func Find(domain string) (Policy, bool) {

	mutex.RLock()
	defer mutex.RUnlock()

	policy, ok := policies[domain]

	return policy, ok
}
