package v1

import (
	"github.com/5krotov/task-resolver-pkg/entity/v1"
)

type CreateTaskRequest struct {
	Name       string `json:"name"`
	Difficulty int    `json:"difficulty"`
}

type CreateTaskResponse struct {
	v1.Task
}
