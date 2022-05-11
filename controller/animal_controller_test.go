package controller_test

import (
	"encoding/json"
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
}
