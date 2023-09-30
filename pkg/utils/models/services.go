package models

type CreateCategory struct {
	Category string `json:"category"`
}

type AddServicesToACategory struct {
	CategoryID  int    `json:"category_id"`
	ServiceName string `json:"service"`
}

type PlaceBid struct {
	WorkID      int     `json:"-"`
	ProID       int     `json:"-"`
	Estimate    float64 `json:"estimate"`
	Description string  `json:"description"`
}
