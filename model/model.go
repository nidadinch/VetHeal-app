package model

type Animal struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Animal_type  string `json:"type"`
	Animal_order int    `json:"animal_order"`
	Image        string `json:"image"`
}

type Symptom struct {
	Id                int    `json:"id"`
	Animal_id         int    `json:"animal_id"`
	Description       string `json:"description"`
	Created_at        string `json:"created_at"`
	Updated_at        string `json:"updated_at"`
	Initial_action_id int    `json:"initial_action_id"`
}

type Action struct {
	Id              int    `json:"id"`
	Actionable_type string `json:"type"`
	Created_at      string `json:"created_at"`
	Updated_at      string `json:"updated_at"`
}

type Question struct {
	Actionable_id int      `json:"actionable_id"`
	Text          string   `json:"text"`
	Options       []Option `json:"options"`
}

type Option struct {
	Text           string `json:"text"`
	Next_action_id int    `json:"next_action_id"`
}

type Result struct {
	Risk_category         RiskCategory `json:"risk_category"`
	Additional_advice     string       `json:"additional_advice"`
	First_aid_text        string       `json:"first_aid_text"`
	Problem_text          string       `json:"problem_text"`
	Travel_advice_text    string       `json:"travel_advice_text"`
	Iframe_first_aid_text string       `json:"iframe_first_aid_text"`
	Iframe_problem_text   string       `json:"iframe_problem_text"`
}

type RiskCategory struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Text_1        string `json:"text_1"`
	Iframe_desc   string `json:"iframe_desc"`
	Iframe_text_1 string `json:"iframe_text_1"`
	Country_id    string `json:"country_id"`
	Created_at    string `json:"created_at"`
	Updated_at    string `json:"updated_at"`
	Rating        string `json:"rating"`
}

type ActionableQuestion struct {
	Id              int      `json:"id"`
	Actionable_type string   `json:"type"`
	Created_at      string   `json:"created_at"`
	Updated_at      string   `json:"updated_at"`
	Question        Question `json:"question"`
}

type ActionableResult struct {
	Id              int    `json:"id"`
	Actionable_type string `json:"type"`
	Created_at      string `json:"created_at"`
	Updated_at      string `json:"updated_at"`
	Result          Result `json:"result"`
}
type Actionable struct {
	Id              int      `json:"id"`
	Actionable_type string   `json:"type"`
	Created_at      string   `json:"created_at"`
	Updated_at      string   `json:"updated_at"`
	Question        Question `json:"question"`
	Result          Result   `json:"result"`
}
type AnimalsResponse []Animal
type SymptomsResponse []Symptom
