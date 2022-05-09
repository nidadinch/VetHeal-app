package repository

import (
	"database/sql"
	"seniorproject-backend/database"
	"seniorproject-backend/model"
)

type Animal struct {
	DB *sql.DB
}

type IAnimal interface {
	GetAnimals() ([]*model.Animal, error)
}

func (a *Animal) GetAnimals() ([]*model.Animal, error) {
	rows, err := a.DB.Query("SELECT * FROM animal")
	database.CheckErr(err)
	var animals []*model.Animal

	for rows.Next() {
		var id int
		var name string
		var animal_type string
		var order int
		var image string

		err = rows.Scan(&id, &name, &animal_type, &order, &image)

		// check errors
		database.CheckErr(err)

		animals = append(animals, &model.Animal{Id: id, Name: name, Animal_type: animal_type, Order: order, Image: image})
	}
	return animals, err
}

func NewAnimalRepository() IAnimal {
	database := database.SetupDB()
	return &Animal{DB: database}
}
