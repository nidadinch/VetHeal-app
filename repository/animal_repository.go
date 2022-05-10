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
	GetActionableType(id string) string
	GetActionableQuestion(id string) (*model.ActionableQuestion, error)
	GetActionableResult(id string) (*model.ActionableResult, error)
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

func (a *Animal) GetActionableType(id string) string {
	var action_type string
	a.DB.QueryRow("SELECT * FROM actionable id = $1", id).Scan(&action_type)

	return action_type
}

func (a *Animal) GetActionableQuestion(id string) (*model.ActionableQuestion, error) {
	rows, err := a.DB.Query("SELECT * FROM symptom WHERE animal_id = $1", id)
	database.CheckErr(err)
	var actionableQuestion *model.ActionableQuestion

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

		actionableQuestion = &model.ActionableQuestion{Id: id}
	}
	return actionableQuestion, err
}

func (a *Animal) GetActionableResult(id string) (*model.ActionableResult, error) {
	var actionableResult *model.ActionableResult
	var riskCategory *model.RiskCategory
	var result *model.Result

	rows, err := a.DB.Query("select actionable.*, result.*, risk_category.name, risk_category.description, risk_category.text_1, risk_category.iframe_desc, risk_category.iframe_text_1,risk_category.rating, risk_category.country_id, risk_category.created_at,risk_category.updated_at from  response  left join result on result.response_id = response.id  left join risk_category on risk_category.id = result.risk_category_id  left join actionable on actionable.id = response.action_id where response.action_id = action_id and action_id = $1", id)
	for rows.Next() {
		var id int
		var actionable_type string
		var created_at string
		var updated_at string
		var response_id int
		var additional_advice string
		var first_aid_text string
		var problem_text string
		var travel_advice_text string
		var iframe_first_aid_text string
		var iframe_problem_text string
		var risk_category_id string
		var name string
		var description string
		var text_1 string
		var iframe_desc string
		var iframe_text_1 string
		var rating string
		var country_id string
		var r_created_at string
		var r_updated_at string

		err = rows.Scan(&id, &created_at, &updated_at, &actionable_type, &response_id,
			&additional_advice, &first_aid_text, &problem_text, &travel_advice_text,
			&iframe_problem_text, &iframe_first_aid_text, &risk_category_id, &name, &description, &text_1,
			&iframe_desc, &iframe_text_1, &rating, &country_id, r_created_at, r_updated_at)

		// check errors
		database.CheckErr(err)
		riskCategory = &model.RiskCategory{Name: name, Description: description, Text_1: text_1, Iframe_desc: iframe_desc,
			Iframe_text_1: iframe_text_1, Country_id: country_id, Rating: rating}
		result = &model.Result{Risk_category: *riskCategory, Additional_advice: additional_advice,
			First_aid_text: first_aid_text, Problem_text: problem_text, Travel_advice_text: travel_advice_text,
			Iframe_first_aid_text: iframe_first_aid_text, Iframe_problem_text: iframe_problem_text}
		actionableResult = &model.ActionableResult{Id: id, Actionable_type: actionable_type,
			Created_at: created_at, Updated_at: updated_at, Result: *result}
	}

	return actionableResult, err
}

func NewAnimalRepository() IAnimal {
	database := database.SetupDB()
	return &Animal{DB: database}
}
