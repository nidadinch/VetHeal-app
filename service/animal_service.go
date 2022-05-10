package service

import (
	"seniorproject-backend/model"
	"seniorproject-backend/repository"
)

type IAnimalService interface {
	Animals() (*model.AnimalsResponse, error)
	Symptoms(animalId string) (*model.SymptomsResponse, error)
	Actionables(id string) (interface{}, error)
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

func (s *AnimalService) Actionables(id string) (interface{}, error) {
	actionableType := s.Repository.GetActionableType(id)

	if actionableType == "Question" {
		actionableQuestion, err := s.Repository.GetActionableQuestion(id)
		return *actionableQuestion, err

	} else {
		actionableResult, err := s.Repository.GetActionableResult(id)
		return *actionableResult, err
	}

	return nil, nil
}

func NewAnimalService(repository repository.IAnimal) IAnimalService {
	return &AnimalService{Repository: repository}
}
