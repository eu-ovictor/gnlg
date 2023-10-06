package relationship

type Usecase interface {
    Add(rel Relationship) error 
}
