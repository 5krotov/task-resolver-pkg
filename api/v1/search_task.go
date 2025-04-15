package v1

import (
	"github.com/5krotov/task-resolver-pkg/entity/v1"
)

type SearchTaskResponse struct {
	Pages int       `json:"pages"`
	Tasks []v1.Task `json:"tasks"`
}
