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

func (u personUsecase) Edit(p person.Person) error {
	return u.repository.Edit(p)
}

func (u personUsecase) Fetch(ID int, name string) ([]person.Person, error) {
	return u.repository.Fetch(ID, name)
}

func (u personUsecase) Delete(ID int) error {
	return u.repository.Delete(ID)
}
