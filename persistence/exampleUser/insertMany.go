package exampleUser

import (
	"template_ddd_service/domain"
)

type InsertEntries []InsertEntry

func (e InsertEntries) Table() string {
	return "entries"
}

func (e InsertEntries) GetFirstItem() any {
	if len(e) == 0 {
		return InsertEntry{}
	}
	return e[0]
}

func (e InsertEntries) GetItems() []any {
	if len(e) == 0 {
		return nil
	}
	var result []any
	for _, entry := range e {
		result = append(result, entry)
	}
	return result
}

func (e InsertEntries) Len() int {
	return len(e)
}

func (e SelectEntries) ToDomain() domain.Users {
	res := make(domain.Users, len(e))
	for i, ee := range e {
		res[i] = domain.User{
			ID: ee.ID,
		}
	}
	return res
}

func ToInsertEntries(e domain.Users) InsertEntries {
	res := make(InsertEntries, len(e))
	for i, ee := range e {
		res[i] = InsertEntry{
			ID: ee.ID,
		}
	}
	return res
}
