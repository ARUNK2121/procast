package models

type CreateCategory struct {
	Category string `json:"category"`
}

type AddServicesToACategory struct {
	CategoryID  int    `json:"category_id"`
	ServiceName string `json:"service"`
}
