package service

import (
	"seniorproject-backend/model"
	"seniorproject-backend/repository"
)

type IAnimalService interface {
	Animals() (*model.AnimalsResponse, error)
	Symptoms(animalId string) (*model.SymptomsResponse, error)
	Actionables(id string) (*model.Actionable, error)
}

type AnimalService struct {
	Repository repository.IAnimal
}

func (s *AnimalService) Animals() (*model.AnimalsResponse, error) {
	animals, err := s.Repository.GetAnimals()
	m := model.AnimalsResponse{}

	for _, v := range animals {
		m = append(m, *v)
	}
	return &m, err
}

func (s *AnimalService) Symptoms(animalId string) (*model.SymptomsResponse, error) {
	sypmtoms, err := s.Repository.GetSymptoms(animalId)
	m := model.SymptomsResponse{}

	for _, v := range sypmtoms {
		m = append(m, *v)
	}
	return &m, err
}

func (s *AnimalService) Actionables(id string) (*model.Actionable, error) {
	// actionable, err := s.Repository.GetActionable()
	// m := model.Actionable{}

	// m = *actionable
	// return &m, err
	return &model.Actionable{}, nil

}

func NewAnimalService(repository repository.IAnimal) IAnimalService {
	return &AnimalService{Repository: repository}
}
