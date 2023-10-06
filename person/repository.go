package person

type Repository interface {
	Add(p Person) error
	Edit(p Person) (int64, error)
	Fetch() ([]Person, error)
    GetById(ID int64) (Person, error)
}