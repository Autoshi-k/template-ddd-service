package exampleUser

import "time"

type SelectEntry struct {
	ID        int       `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	DeletedAt time.Time `db:"deleted_at"`
}

func (SelectEntry) Table() string {
	return "users"
}

type SelectEntryCdn map[SelectEntryColumnName]any

type SelectEntryColumnName string

const (
	ID SelectEntryColumnName = "id"
)
