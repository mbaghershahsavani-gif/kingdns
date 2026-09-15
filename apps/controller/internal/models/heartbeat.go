package models

import "time"

type NodeHeartbeat struct {
	ID        uint      `gorm:"primaryKey"`
	NodeID    uint
	Latency   int
	CPU       int
	Memory    int
	Status    string
	CreatedAt time.Time
}
