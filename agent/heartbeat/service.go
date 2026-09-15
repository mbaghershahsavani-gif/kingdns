package heartbeat

import (
    "time"
)

type Report struct {
    NodeID int `json:"node_id"`
    Status string `json:"status"`
    CPU float64 `json:"cpu"`
    Memory float64 `json:"memory"`
}

func Start(interval time.Duration, send func(Report)) {
    ticker := time.NewTicker(interval)

    go func() {
        for range ticker.C {
            send(Report{
                Status: "online",
            })
        }
    }()
}
