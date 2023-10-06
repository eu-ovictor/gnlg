package person

type Usecase interface {
	Add(p Person) error
	Edit(p Person) (int64, error)
	Fetch() ([]Person, error)
}
