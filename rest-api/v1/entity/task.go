package entity

type Task struct {
	Id            int64    `json:"id"`
	Name          string   `json:"name"`
	Difficulty    int      `json:"difficulty"`
	StatusHistory []Status `json:"status_history"`
}
