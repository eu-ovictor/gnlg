package person

type Usecase interface {
	Add(p Person) error
	Edit(p Person) error
	Fetch(ID int, name string) ([]Person, error)
}
