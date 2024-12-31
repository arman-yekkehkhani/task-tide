package chore

type ID int64

type Chore struct {
	ID
	Title       string `json:"title"`
	Description string `json:"description"`
}
