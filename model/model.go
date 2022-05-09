package model

type Animal struct {
	id          int    `json:"id"`
	name        string `json:"name"`
	animal_type string `json:"type"`
	order       int    `json:"order"`
	image       string `json:"image"`
}

type Symptom struct {
	id                int    `json:"id"`
	animal_id         int    `json:"animal_id"`
	desc              string `json:"desc"`
	initial_action_id int    `json:"initial_action_id"`
}

type Actionable struct {
	id              int    `json:"id"`
	actionable_type string `json:"type"`
	question        string `json:"question"`
	result          string `json:"result"`
}

type Question struct {
	actionable_id int      `json:"actionable_id"`
	text          string   `json:"text"`
	options       []Option `json:"options"`
}

type Option struct {
	text           string `json:"text"`
	next_action_id int    `json:"next_action_id"`
}

type Result struct {
	risk_category         RiskCategory `json:"risk_category"`
	additional_advice     string       `json:"additional_advice"`
	first_aid_text        string       `json:"first_aid_text"`
	problem_text          string       `json:"problem_text"`
	travel_advice_text    string       `json:"travel_advice_text"`
	iframe_first_aid_text string       `json:"iframe_first_aid_text"`
	iframe_problem_text   string       `json:"iframe_problem_text"`
}

type RiskCategory struct {
	name          string `json:"name"`
	desc          string `json:"desc"`
	text_1        string `json:"text_1"`
	iframe_desc   string `json:"iframe_desc"`
	iframe_text_1 string `json:"iframe_text_1"`
	country_id    string `json:"country_id"`
	created_at    string `json:"created_at"`
	updated_at    string `json:"updated_at"`
	rating        string `json:"rating"`
}

type AnimalsResponse []Animal
type SymptomsResponse []Symptom
