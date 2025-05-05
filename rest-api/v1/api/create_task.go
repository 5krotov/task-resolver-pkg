package api

import (
	"github.com/5krotov/task-resolver-pkg/rest-api/v1/entity"
)

type CreateTaskRequest struct {
	Name       string `json:"name"`
	Difficulty int    `json:"difficulty"`
}

type CreateTaskResponse struct {
	entity.Task
}
