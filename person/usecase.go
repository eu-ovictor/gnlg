package person

type PersonUsecase interface {
    Add(p Person) error 
    Edit(p Person) (int64, error)
    Fetch() ([]Person, error)
}
