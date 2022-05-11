package repository

import (
	"database/sql"
	"seniorproject-backend/database"
	"seniorproject-backend/model"

	"gopkg.in/guregu/null.v3"
)

type Animal struct {
	DB *sql.DB
}

type IAnimal interface {
	GetAnimals() ([]*model.Animal, error)
	GetSymptoms(animalId string) ([]*model.Symptom, error)
	GetActionableType(id string) string
	GetQuestionActionable(id string) (*model.QuestionActionable, error)
	GetResultActionable(id string) (*model.ResultActionable, error)
	GetActionableOptions(id string) []*model.Option
}

func (a *Animal) GetAnimals() ([]*model.Animal, error) {
	rows, err := a.DB.Query("SELECT * FROM animal")
	database.CheckErr(err)
	var animals []*model.Animal

	for rows.Next() {
		var id int
		var name null.String
		var animal_type null.String
		var image null.String
		var animal_order null.Int

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
		var animal_id null.Int
		var description null.String
		var created_at null.String
		var updated_at null.String
		var initial_action_id null.Int

		err = rows.Scan(&id, &animal_id, &description, &created_at, &updated_at, &initial_action_id)

		// check errors
		database.CheckErr(err)

		symptoms = append(symptoms, &model.Symptom{Id: id, Animal_id: animal_id, Description: description, Created_at: created_at, Updated_at: updated_at, Initial_action_id: initial_action_id})
	}
	return symptoms, err
}

func (a *Animal) GetActionableType(id string) string {
	var action_type string
	a.DB.QueryRow("SELECT action_type FROM actionable WHERE id = $1", id).Scan(&action_type)

	return action_type
}
func (a *Animal) GetActionableOptions(id string) []*model.Option {
	optionsQuery, err := a.DB.Query("select option.text as option_text, option.next_action_id, option.response_id from response	left join option on option.response_id = response.id left join question on question.actionable_id = response.action_id where response.action_id = response.action_id and  question.actionable_id = response.action_id and question.actionable_id= $1", id)
	database.CheckErr(err)
	var options []*model.Option
	for optionsQuery.Next() {
		var option_text null.String
		var next_action_id null.Int
		var response_id null.Int

		err = optionsQuery.Scan(&option_text, &next_action_id, &response_id)

		// check errors
		database.CheckErr(err)

		options = append(options, &model.Option{Text: option_text, Next_action_id: next_action_id, Response_id: response_id})
	}
	return options
}

func (a *Animal) GetQuestionActionable(id string) (*model.QuestionActionable, error) {
	options := a.GetActionableOptions(id)
	rows, err := a.DB.Query("select actionable.*, question.* from response left join option on option.response_id = response.id left join question on question.actionable_id = response.action_id left join actionable on question.actionable_id = actionable.id where response.action_id = response.action_id and  question.actionable_id = response.action_id and question.actionable_id = $1", id)
	database.CheckErr(err)
	var actionableQuestion *model.QuestionActionable
	var question *model.Question

	for rows.Next() {
		var id int
		var created_at null.String
		var updated_at null.String
		var action_type null.String
		var actionable_id int
		var text null.String

		err = rows.Scan(&id, &created_at, &updated_at, &action_type, &actionable_id, &text)

		// check errors
		database.CheckErr(err)
		question = &model.Question{Text: text, Actionable_id: actionable_id, Options: options}

		actionableQuestion = &model.QuestionActionable{Id: id, Created_at: created_at, Updated_at: updated_at, Actionable_type: action_type,
			Question: *question}
	}
	return actionableQuestion, err
}

func (a *Animal) GetResultActionable(id string) (*model.ResultActionable, error) {
	var actionableResult *model.ResultActionable
	var riskCategory *model.RiskCategory
	var result *model.Result

	rows, err := a.DB.Query("select actionable.*, result.*, risk_category.name, risk_category.description, risk_category.text_1, risk_category.iframe_desc, risk_category.iframe_text_1,risk_category.rating, risk_category.country_id, risk_category.created_at,risk_category.updated_at from  response  left join result on result.response_id = response.id  left join risk_category on risk_category.id = result.risk_category_id  left join actionable on actionable.id = response.action_id where response.action_id = action_id and action_id = $1", id)
	for rows.Next() {
		var id int
		var action_type null.String
		var created_at null.String
		var updated_at null.String
		var response_id null.Int
		var additional_advice null.String
		var first_aid_text null.String
		var problem_text null.String
		var travel_advice_text null.String
		var iframe_first_aid_text null.String
		var iframe_problem_text null.String
		var risk_category_id null.String
		var name null.String
		var description null.String
		var text_1 null.String
		var iframe_desc null.String
		var iframe_text_1 null.String
		var rating null.String
		var country_id null.String
		var r_created_at null.String
		var r_updated_at null.String

		err = rows.Scan(&id, &created_at, &updated_at, &action_type, &response_id,
			&additional_advice, &first_aid_text, &problem_text, &travel_advice_text,
			&iframe_first_aid_text, &iframe_problem_text, &risk_category_id, &name, &description, &text_1,
			&iframe_desc, &iframe_text_1, &rating, &country_id, &r_created_at, &r_updated_at)

		// check errors
		database.CheckErr(err)
		riskCategory = &model.RiskCategory{Name: name, Description: description, Text_1: text_1, Iframe_desc: iframe_desc,
			Iframe_text_1: iframe_text_1, Country_id: country_id, Rating: rating}
		result = &model.Result{Risk_category: *riskCategory, Additional_advice: additional_advice,
			First_aid_text: first_aid_text, Problem_text: problem_text, Travel_advice_text: travel_advice_text,
			Iframe_first_aid_text: iframe_first_aid_text, Iframe_problem_text: iframe_problem_text}
		actionableResult = &model.ResultActionable{Id: id, Actionable_type: action_type,
			Created_at: created_at, Updated_at: updated_at, Result: *result}
	}

	return actionableResult, err
}

func NewAnimalRepository() IAnimal {
	database := database.SetupDB()
	return &Animal{DB: database}
}
