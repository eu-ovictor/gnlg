package usecase

import (
	"github.com/eu-ovictor/gnlg/person"
)

type personUsecase struct {
	repository person.Repository
}

func NewPersonUsecase(repository person.Repository) personUsecase {
	return personUsecase{
		repository,
	}
}

func (u personUsecase) Add(p person.Person) error {
	return u.repository.Add(p)
}

func (u personUsecase) Edit(p person.Person) (int64, error) {
	return u.repository.Edit(p)
}

func (u personUsecase) Fetch() ([]person.Person, error) {
	return u.repository.Fetch()
}
