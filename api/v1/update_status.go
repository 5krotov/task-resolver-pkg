package v1

import (
	"github.com/5krotov/task-resolver-pkg/entity/v1"
)

type UpdateStatusRequest struct {
	Id     int64     `json:"id"`
	Status v1.Status `json:"status"`
}
