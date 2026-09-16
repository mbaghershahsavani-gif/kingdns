package cluster

import (
	"encoding/json"
	"net/http"
	"time"
)

func SyncPeers() {

	peers := GetPeers()

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	for _, peer := range peers {

		resp, err := client.Get(
			peer + "/cluster/sync",
		)

		if err != nil {
			continue
		}

		var scores []Score

		err =
			json.NewDecoder(resp.Body).
				Decode(&scores)

		resp.Body.Close()

		if err != nil {
			continue
		}

		for _, score := range scores {

			RegisterScore(score)
		}
	}
}
