package api

import (
	"github.com/5krotov/task-resolver-pkg/rest-api/v1/entity"
)

type SearchTaskResponse struct {
	Pages int           `json:"pages"`
	Tasks []entity.Task `json:"tasks"`
}
