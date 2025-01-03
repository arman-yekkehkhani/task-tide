package base

import "time"

type ID int64

type BaseEntity struct {
	ID
	createdAt time.Time
	updatedAt time.Time
	IsDeleted bool
}
