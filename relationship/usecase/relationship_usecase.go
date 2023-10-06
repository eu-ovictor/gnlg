package usecase

import "github.com/eu-ovictor/gnlg/relationship"

type relationshipUsecase struct {
	repository relationship.Repository
}

func NewRelationshipUsecase(repository relationship.Repository) relationshipUsecase {
	return relationshipUsecase{repository}
}

func (u relationshipUsecase) Add(rel relationship.Relationship) error {
    return u.repository.Add(rel)
}
