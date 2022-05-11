package service

import (
	"fmt"
	"seniorproject-backend/model"
	"seniorproject-backend/repository"
)

type IAnimalService interface {
	Animals() (*model.AnimalsResponse, error)
	Symptoms(animalId string) (*model.SymptomsResponse, error)
	GetActionable(id string) (interface{}, error)
	ActionableQuestion(id string) (*model.QuestionActionable, error)
	ActionableResult(id string) (*model.ResultActionable, error)
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

func (s *AnimalService) GetActionable(id string) (interface{}, error) {
	actionableType := s.Repository.GetActionableType(id)
	fmt.Println(actionableType)

	if actionableType == "Question" {
		fmt.Println("id: ", id)
		fmt.Println("Question")
		result, err := s.ActionableQuestion(id)
		return result, err

	} else {
		fmt.Println("Result")
		fmt.Println("id: ", id)
		result, err := s.ActionableResult(id)
		return result, err

	}
}

func (s *AnimalService) ActionableQuestion(id string) (*model.QuestionActionable, error) {
	actionableQuestion, err := s.Repository.GetQuestionActionable(id)

	m := model.QuestionActionable{}
	m = *actionableQuestion

	return &m, err
}

func (s *AnimalService) ActionableResult(id string) (*model.ResultActionable, error) {
	actionableResult, err := s.Repository.GetResultActionable(id)

	m := model.ResultActionable{}
	m = *actionableResult

	return &m, err
}

func NewAnimalService(repository repository.IAnimal) IAnimalService {
	return &AnimalService{Repository: repository}
}
