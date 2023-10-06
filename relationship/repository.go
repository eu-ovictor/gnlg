package relationship

type Repository interface {
	Add(r Relationship) error
}
