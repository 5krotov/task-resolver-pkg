package entity

import "time"

type Status struct {
	Status    int       `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

const (
	STATUS_CREATED    = 0
	STATUS_PENDING    = 1
	STATUS_INPROGRESS = 2
	STATUS_DONE       = 3
)
