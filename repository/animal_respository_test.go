package repository_test

import (
	"fmt"
	"seniorproject-backend/model"
	"seniorproject-backend/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
)

func Test_GetAllAnimals(t *testing.T) {
	t.Run("should return all animals correctly", func(t *testing.T) {
		animalRepo, animals := prepareForTest()
		want := animals
		got, err := animalRepo.GetAnimals()
		assert.Nil(t, err)
		assert.Equal(t, got, want)

	})

	t.Run("should return empty animal if DB is empty", func(t *testing.T) {
		animals := []*model.Animal{}

		animalRepo := repository.NewAnimalRepository()

		want := animals
		got, err := animalRepo.GetAnimals()
		assert.Nil(t, err)
		assert.Equal(t, got, want)
	})
}

func prepareForTest() (repository.IAnimal, []*model.Animal) {
	animals := []*model.Animal{
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
	// Create repository
	animalRepo := repository.NewAnimalRepository()
	// Append items to repository
	for _, s := range animals {
		fmt.Println(s)
	}

	return animalRepo, animals
}
