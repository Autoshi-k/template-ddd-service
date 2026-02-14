package exampleUser

import "template_ddd_service/domain"

type InsertEntry struct {
	ID int `db:"id"`
}

func (InsertEntry) Table() string {
	return "users"
}

func ToInsertEntry(v domain.User) InsertEntry {
	return InsertEntry{
		ID: v.ID,
	}
}
