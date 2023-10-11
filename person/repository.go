package person

type Repository interface {
	Add(p Person) error
	Edit(p Person) error
	Fetch(ID int, name string) ([]Person, error)
	Delete(ID int) error
}
