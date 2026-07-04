package model

import "time"

type WorkerStatus string

const (
	WorkerStatusAlive WorkerStatus = "alive"
	WorkerStatusDead  WorkerStatus = "dead"
)

type Workers struct {
	ID            string       `json:"id" db:"id"`
	Hostname      string       `json:"hostname" db:"hostname"`
	Status        WorkerStatus `json:"status" db:"status"`
	Capabilities  []string     `json:"capabilities" db:"capabilities"`
	LastHeartbeat time.Time    `json:"last_heartbeat" db:"last_heartbeat"`
	RegisterdAt   time.Time    `json:"registered_at" db:"registered_at"`
}
