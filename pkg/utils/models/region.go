package models

type AddNewState struct {
	State string `json:"state"`
}

type AddNewDistrict struct {
	StateID  int    `json:"state_id"`
	District string `json:"district"`
}
