package usecase

import "github.com/eu-ovictor/gnlg/relationship"

type relationshipUsecase struct {
	repository relationship.Repository
}

func NewRelationshipUsecase(repository relationship.Repository) relationshipUsecase {
	return relationshipUsecase{repository}
}

func (u relationshipUsecase) Add(members relationship.Members) error {
	return u.repository.Add(members)
}

func (u relationshipUsecase) FetchByID(ID int64) (relationship.Relationships, error) {
	members, err := u.repository.FetchByID(ID)
	if err != nil {
		return nil, err
	}

	relationships := make(relationship.Relationships, len(members))

	for _, member := range members {
		rel := relationship.Relationship{Member: member.Parent, Kinship: relationship.Parent}

		relationships[member.Child] = append(relationships[member.Child], rel)
	}

	return relationships, nil
}
