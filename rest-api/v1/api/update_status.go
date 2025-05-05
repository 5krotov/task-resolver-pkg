package api

import (
	"github.com/5krotov/task-resolver-pkg/rest-api/v1/entity"
)

type UpdateStatusRequest struct {
	Id     int64         `json:"id"`
	Status entity.Status `json:"status"`
}
