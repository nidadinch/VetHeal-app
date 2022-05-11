package model

import (
	"gopkg.in/guregu/null.v3"
)

type Animal struct {
	Id           int         `json:"id"`
	Name         null.String `json:"name"`
	Animal_type  null.String `json:"type"`
	Animal_order null.Int    `json:"animal_order"`
	Image        null.String `json:"image"`
}

type Symptom struct {
	Id                int         `json:"id"`
	Animal_id         null.Int    `json:"animal_id"`
	Description       null.String `json:"description"`
	Created_at        null.String `json:"created_at"`
	Updated_at        null.String `json:"updated_at"`
	Initial_action_id null.Int    `json:"initial_action_id"`
}

type Action struct {
	Id              int         `json:"id"`
	Actionable_type null.String `json:"type"`
	Created_at      null.String `json:"created_at"`
	Updated_at      null.String `json:"updated_at"`
}

type Question struct {
	Actionable_id int         `json:"actionable_id"`
	Text          null.String `json:"text"`
	Options       []*Option   `json:"options"`
}

type Option struct {
	Text           null.String `json:"text"`
	Next_action_id null.Int    `json:"next_action_id"`
	Response_id    null.Int    `json:"response_id"`
}

type Result struct {
	Risk_category         RiskCategory `json:"risk_category"`
	Additional_advice     null.String  `json:"additional_advice"`
	First_aid_text        null.String  `json:"first_aid_text"`
	Problem_text          null.String  `json:"problem_text"`
	Travel_advice_text    null.String  `json:"travel_advice_text"`
	Iframe_first_aid_text null.String  `json:"iframe_first_aid_text"`
	Iframe_problem_text   null.String  `json:"iframe_problem_text"`
}

type RiskCategory struct {
	Name          null.String `json:"name"`
	Description   null.String `json:"description"`
	Text_1        null.String `json:"text_1"`
	Iframe_desc   null.String `json:"iframe_desc"`
	Iframe_text_1 null.String `json:"iframe_text_1"`
	Country_id    null.String `json:"country_id"`
	Created_at    null.String `json:"created_at"`
	Updated_at    null.String `json:"updated_at"`
	Rating        null.String `json:"rating"`
}

type QuestionActionable struct {
	Id              int         `json:"id"`
	Actionable_type null.String `json:"type"`
	Created_at      null.String `json:"created_at"`
	Updated_at      null.String `json:"updated_at"`
	Question        Question    `json:"question"`
}

type ResultActionable struct {
	Id              int         `json:"id"`
	Actionable_type null.String `json:"type"`
	Created_at      null.String `json:"created_at"`
	Updated_at      null.String `json:"updated_at"`
	Result          Result      `json:"result"`
}

type AnimalsResponse []Animal
type SymptomsResponse []Symptom
