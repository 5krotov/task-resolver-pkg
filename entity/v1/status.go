package v1

import "time"

type Status struct {
	Status    int       `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}
