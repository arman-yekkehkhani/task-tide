package chore

import . "github.com/arman-yekkehkhani/task-tide/internal/model/base"

type Status string

const (
	PENDING     Status = "pending"
	IN_PROGRESS Status = "in_progress"
	DONE        Status = "done"
)

type Chore struct {
	ID          ID
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status
}
