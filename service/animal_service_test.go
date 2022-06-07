package service_test

import (
	"seniorproject-backend/mock"
	"seniorproject-backend/model"
	"seniorproject-backend/service"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
)

func TestService_GetAnimals(t *testing.T) {
	t.Run("should return animals correctly", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockData := mock.NewMockIAnimal(mockController)

		mockAnimals := []*model.Animal{{
			Id:           1,
			Name:         null.NewString("Dog", true),
			Animal_type:  null.NewString("companion", true),
			Animal_order: null.NewInt(1, true),
			Image:        null.NewString("https://media.nature.com/lw800/magazine-assets/d41586-022-00209-0/d41586-022-00209-0_20071828.jpg", true),
		},
			{
				Id:           2,
				Name:         null.NewString("Cats", true),
				Animal_type:  null.NewString("companion", true),
				Animal_order: null.NewInt(2, true),
				Image:        null.NewString("https://www.inquirer.com/resizer/A12f2p2Ga06aYaV4Wm8R69zGZ6Y=/4x0:5568x3712/760x507/filters:format(webp)/cloudfront-us-east-1.images.arcpublishing.com/pmn/EDNMWLP6FZFFNKSNE76TI5K7RQ.jpg", true),
			}}
		mockData.EXPECT().GetAnimals().Return(mockAnimals, nil).Times(1)

		AnimalService := service.NewAnimalService(mockData)
		animals, err := AnimalService.Animals()
		AnimalsResponse := &model.AnimalsResponse{
			{
				Id:           1,
				Name:         null.NewString("Dog", true),
				Animal_type:  null.NewString("companion", true),
				Animal_order: null.NewInt(1, true),
				Image:        null.NewString("https://media.nature.com/lw800/magazine-assets/d41586-022-00209-0/d41586-022-00209-0_20071828.jpg", true),
			},
			{
				Id:           2,
				Name:         null.NewString("Cats", true),
				Animal_type:  null.NewString("companion", true),
				Animal_order: null.NewInt(2, true),
				Image:        null.NewString("https://www.inquirer.com/resizer/A12f2p2Ga06aYaV4Wm8R69zGZ6Y=/4x0:5568x3712/760x507/filters:format(webp)/cloudfront-us-east-1.images.arcpublishing.com/pmn/EDNMWLP6FZFFNKSNE76TI5K7RQ.jpg", true),
			},
		}
		assert.Equal(t, AnimalsResponse, animals)
		assert.Nil(t, err)
	})
}

func TestService_GetSymptoms(t *testing.T) {
	t.Run("should return symptoms correctly", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockData := mock.NewMockIAnimal(mockController)

		mockSymptoms := []*model.Symptom{{
			Id:                31,
			Animal_id:         null.NewInt(1, true),
			Description:       null.NewString("Collapsed", true),
			Created_at:        null.NewString("2019-01-03T09:48:40.000000Z", true),
			Updated_at:        null.NewString("2019-01-03T09:48:40.000000Z", true),
			Initial_action_id: null.NewInt(4456, true),
		},
			{
				Id:                32,
				Animal_id:         null.NewInt(1, true),
				Description:       null.NewString("Diarrhoea", true),
				Created_at:        null.NewString("2019-01-03T09:48:40.000000Z", true),
				Updated_at:        null.NewString("2019-01-03T09:48:40.000000Z", true),
				Initial_action_id: null.NewInt(726, true),
			}}

		animalId := "1"
		mockData.EXPECT().GetSymptoms(animalId).Return(mockSymptoms, nil).Times(1)

		AnimalService := service.NewAnimalService(mockData)
		symptoms, err := AnimalService.Symptoms(animalId)
		SymptomsResponse := &model.SymptomsResponse{
			{
				Id:                31,
				Animal_id:         null.NewInt(1, true),
				Description:       null.NewString("Collapsed", true),
				Created_at:        null.NewString("2019-01-03T09:48:40.000000Z", true),
				Updated_at:        null.NewString("2019-01-03T09:48:40.000000Z", true),
				Initial_action_id: null.NewInt(4456, true),
			},
			{
				Id:                32,
				Animal_id:         null.NewInt(1, true),
				Description:       null.NewString("Diarrhoea", true),
				Created_at:        null.NewString("2019-01-03T09:48:40.000000Z", true),
				Updated_at:        null.NewString("2019-01-03T09:48:40.000000Z", true),
				Initial_action_id: null.NewInt(726, true),
			},
		}
		assert.Equal(t, SymptomsResponse, symptoms)
		assert.Nil(t, err)
	})
}
