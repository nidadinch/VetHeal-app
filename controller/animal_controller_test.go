package controller_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"seniorproject-backend/controller"
	"seniorproject-backend/mock"
	"seniorproject-backend/model"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
)

func Test_GetAllAnimals(t *testing.T) {
	t.Run("should return all animals correctly", func(t *testing.T) {

		service := mock.NewMockIAnimalService(gomock.NewController(t))

		serviceReturn := &model.AnimalsResponse{
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
			{
				Id:           4,
				Name:         null.NewString("Rabbis", true),
				Animal_type:  null.NewString("companion", true),
				Animal_order: null.NewInt(3, true),
				Image:        null.NewString("https://www.caldervets.co.uk/images/pet-care/rabbits/rabbits-dental-care-for-rabbits-01.jpg", true),
			},
		}

		service.EXPECT().Animals().Return(serviceReturn, nil)
		controller := controller.NewAnimalController(service)

		r := httptest.NewRequest(http.MethodGet, "/animals", nil)
		w := httptest.NewRecorder()

		controller.GetAnimals(w, r)

		actual := serviceReturn
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, serviceReturn, actual)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
		assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("content-type"))
	})
	t.Run("should return error if animal doesn't exist", func(t *testing.T) {
		service := mock.NewMockIAnimalService(gomock.NewController(t))
		serviceErr := errors.New("test err")

		service.EXPECT().Animals().Return(nil, serviceErr)
		controller := controller.NewAnimalController(service)

		r := httptest.NewRequest(http.MethodGet, "/animals", nil)
		w := httptest.NewRecorder()
		controller.GetAnimals(w, r)

		actual := &model.AnimalsResponse{}
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
		assert.Equal(t, w.Body.String(), serviceErr.Error())
	})
}
func Test_GetAllSymptoms(t *testing.T) {
	t.Run("should return symptoms correctly", func(t *testing.T) {
		service := mock.NewMockIAnimalService(gomock.NewController(t))

		serviceReturn := &model.SymptomsResponse{
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
		animalId := "1"
		service.EXPECT().Symptoms(animalId).Return(serviceReturn, nil)
		controller := controller.NewAnimalController(service)

		r := httptest.NewRequest(http.MethodGet, "/animal/1/symptoms", nil)
		w := httptest.NewRecorder()

		controller.GetSymptoms(w, r)

		actual := serviceReturn
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, serviceReturn, actual)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
		assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("content-type"))
	})

	t.Run("should return error if symptom doesn't exist", func(t *testing.T) {
		service := mock.NewMockIAnimalService(gomock.NewController(t))
		serviceErr := errors.New("test err")
		animalId := "1"
		service.EXPECT().Symptoms(animalId).Return(nil, serviceErr)
		controller := controller.NewAnimalController(service)

		r := httptest.NewRequest(http.MethodGet, "/animal/1/symptoms", nil)
		w := httptest.NewRecorder()
		controller.GetSymptoms(w, r)

		actual := &model.SymptomsResponse{}
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
		assert.Equal(t, w.Body.String(), serviceErr.Error())
	})
}

func Test_GetResult(t *testing.T) {
	t.Run("should return actionable correctly", func(t *testing.T) {
		service := mock.NewMockIAnimalService(gomock.NewController(t))
		option1 := &model.Option{
			Text:           null.NewString("no", true),
			Next_action_id: null.NewInt(3247, true),
			Response_id:    null.NewInt(4909, true),
		}
		option2 := &model.Option{
			Text:           null.NewString("Yes", true),
			Next_action_id: null.NewInt(3248, true),
			Response_id:    null.NewInt(4911, true),
		}
		question := &model.Question{
			Actionable_id: 3246,
			Text:          null.NewString("Is she comfortable?", true),
			Options:       []*model.Option{option1, option2},
		}
		serviceReturn := &model.QuestionActionable{

			Id:              31,
			Actionable_type: null.NewString("Question", true),
			Created_at:      null.NewString("2019-01-03T09:48:40.000000Z", true),
			Updated_at:      null.NewString("2019-01-03T09:48:40.000000Z", true),
			Question:        *question,
		}
		actionableId := "3246"
		service.EXPECT().GetActionable(actionableId).Return(serviceReturn, nil)
		controller := controller.NewAnimalController(service)

		r := httptest.NewRequest(http.MethodGet, "/action/3246", nil)
		w := httptest.NewRecorder()

		controller.GetActionable(w, r)

		actual := serviceReturn
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, serviceReturn, actual)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
		assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("content-type"))
	})

	t.Run("should return error if actionable doesn't exist", func(t *testing.T) {
		service := mock.NewMockIAnimalService(gomock.NewController(t))
		serviceErr := errors.New("test err")
		actionableId := "3246"
		service.EXPECT().GetActionable(actionableId).Return(nil, serviceErr)
		controller := controller.NewAnimalController(service)

		r := httptest.NewRequest(http.MethodGet, "/action/3246", nil)
		w := httptest.NewRecorder()
		controller.GetActionable(w, r)

		actual := &model.QuestionActionable{}
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
		assert.Equal(t, w.Body.String(), serviceErr.Error())
	})
}
