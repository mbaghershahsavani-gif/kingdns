package cluster

import (
	"net"
	"time"
)

func CheckDNS(node Node) int {

	start := time.Now()

	conn, err := net.DialTimeout(
		"udp",
		node.IP+":53",
		3*time.Second,
	)

	if err != nil {
		return 0
	}

	conn.Close()

	latency := time.Since(start)

	if latency < 50*time.Millisecond {
		return 30
	}

	if latency < 150*time.Millisecond {
		return 20
	}

	return 10
}

func CheckHeartbeat(hb Heartbeat) int {

	if time.Since(hb.LastSeen) < time.Minute {
		return 40
	}

	return 0
}

func CheckLatency(node Node) int {

	start := time.Now()

	conn, err := net.DialTimeout(
		"tcp",
		node.IP+":8081",
		3*time.Second,
	)

	if err != nil {
		return 0
	}

	conn.Close()

	delay := time.Since(start)

	if delay < 50*time.Millisecond {
		return 10
	}

	if delay < 150*time.Millisecond {
		return 5
	}

	return 1
}
