package relationship

type Usecase interface {
	Add(members Members) error
    FetchByID(ID int64) (Relationships, error)
}
