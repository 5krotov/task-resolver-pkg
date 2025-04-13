package v1

type StartTaskRequest struct {
	Id         int64 `json:"id"`
	Difficulty int   `json:"difficulty"`
}
