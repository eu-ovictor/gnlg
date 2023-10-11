package relationship

type Repository interface {
	Add(members Members) error
    FetchByID(ID int64) ([]NamedMembers, error)
}
