package chore

type ID int64

type Status string

const (
	PENDING     Status = "pending"
	IN_PROGRESS Status = "in_progress"
	DONE        Status = "done"
)

type Chore struct {
	ID
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status
}
