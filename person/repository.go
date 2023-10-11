package person

type Repository interface {
	Add(p Person) error
	Edit(p Person) (int64, error)
	Fetch(ID int, name string) ([]Person, error)
}
