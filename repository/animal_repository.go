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
	GetSymptoms(animalId string) ([]*model.Symptom, error)
}

func (a *Animal) GetAnimals() ([]*model.Animal, error) {
	rows, err := a.DB.Query("SELECT * FROM animal")
	database.CheckErr(err)
	var animals []*model.Animal

	for rows.Next() {
		var id int
		var name string
		var animal_type string
		var image string
		var animal_order int

		err = rows.Scan(&id, &name, &animal_type, &image, &animal_order)

		// check errors
		database.CheckErr(err)

		animals = append(animals, &model.Animal{Id: id, Name: name, Animal_type: animal_type, Animal_order: animal_order, Image: image})
	}
	return animals, err
}

func (a *Animal) GetSymptoms(animalId string) ([]*model.Symptom, error) {

	rows, err := a.DB.Query("SELECT * FROM symptom WHERE animal_id = $1", animalId)
	database.CheckErr(err)
	var symptoms []*model.Symptom

	for rows.Next() {
		var id int
		var animal_id int
		var description string
		var created_at string
		var updated_at string
		var initial_action_id int

		err = rows.Scan(&id, &animal_id, &description, &created_at, &updated_at, &initial_action_id)

		// check errors
		database.CheckErr(err)

		symptoms = append(symptoms, &model.Symptom{Id: id, Animal_id: animal_id, Description: description, Created_at: created_at, Updated_at: updated_at, Initial_action_id: initial_action_id})
	}
	return symptoms, err
}

func NewAnimalRepository() IAnimal {
	database := database.SetupDB()
	return &Animal{DB: database}
}
