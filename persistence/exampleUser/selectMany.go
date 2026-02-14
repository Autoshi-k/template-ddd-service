package exampleUser

type SelectEntries []SelectEntry

func (SelectEntries) Table() string {
	return "users"
}
