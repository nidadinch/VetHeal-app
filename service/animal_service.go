package service

import (
	"seniorproject-backend/model"
	"seniorproject-backend/repository"
)

type IAnimalService interface {
	Animals() (*model.AnimalsResponse, error)
	Symptoms() (*model.SymptomsResponse, error)
	Actionables() (*model.Actionable, error)
}

type AnimalService struct {
	Repository repository.IAnimal
}

func (s *AnimalService) Animals() (*model.AnimalsResponse, error) {
	animals, err := s.Repository.GetAllAnimals()
	m := model.AnimalsResponse{}

	for _, v := range animals {
		m = append(m, *v)
	}
	return &m, err
}

func (s *AnimalService) Symptoms() (*model.SymptomsResponse, error) {
	sypmtoms, err := s.Repository.GetAllSypmtoms()
	m := model.SymptomsResponse{}

	for _, v := range sypmtoms {
		m = append(m, *v)
	}
	return &m, err
}

func (s *AnimalService) Actionables() (*model.Actionable, error) {
	actionable, err := s.Repository.GetActionable()
	m := model.Actionable{}

	m = *actionable
	return &m, err
}

func NewAnimalService(repository repository.IAnimal) IAnimalService {
	return &AnimalService{Repository: repository}
}
