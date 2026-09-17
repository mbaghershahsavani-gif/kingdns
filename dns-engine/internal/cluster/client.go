package cluster

import (
	"encoding/json"
	"log"
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

			log.Printf(
				"Peer sync failed %s: %v",
				peer,
				err,
			)

			continue
		}

		var scores []Score

		err = json.NewDecoder(resp.Body).Decode(&scores)

		resp.Body.Close()

		if err != nil {

			log.Printf(
				"Peer decode failed %s: %v",
				peer,
				err,
			)

			continue
		}

		for _, score := range scores {

			if score.LastSeen.IsZero() {

				log.Printf(
					"Ignoring remote node with empty last_seen: %s",
					score.NodeID,
				)

				continue
			}

			RegisterScore(score)

			log.Printf(
				"Synced remote node: %s region=%s last_seen=%s",
				score.NodeID,
				score.Region,
				score.LastSeen,
			)
		}
	}

	log.Printf(
		"Cluster peer synchronization completed",
	)
}
