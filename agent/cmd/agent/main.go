package main

import (
    "log"
    "time"

    "github.com/mbaghershahsavani-gif/kingdns-agent/heartbeat"
)

func main() {
    log.Println("KingDNS Agent v1.1 starting")

    heartbeat.Start(
        30*time.Second,
        func(report heartbeat.Report) {
            log.Println("heartbeat sent", report.Status)
        },
    )

    select {}
}
