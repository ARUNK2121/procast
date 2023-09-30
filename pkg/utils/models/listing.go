package models

type UserDetails struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email" `
	Phone     string `json:"phone" `
	IsBlocked bool   `json:"is_blocked"`
}

type ProviderDetails struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Verified bool   `json:"verified"`
}

type WorkDetails struct {
	ID         int      `json:"id"`
	Street     string   `json:"street"`
	District   string   `json:"district" `
	State      string   `json:"state" `
	Profession string   `json:"profession"`
	User       string   `json:"user_name"`
	Provider   string   `json:"provider"`
	Images     []string `json:"images" `
	WorkStatus string   `json:"work_status" `
}

type MinWorkDetails struct {
	ID         int    `json:"id"`
	Street     string `json:"street"`
	District   string `json:"district" `
	State      string `json:"state" `
	Profession string `json:"profession"`
	User       string `json:"user_name"`
	WorkStatus string `json:"work_status" `
}

type GetServices struct {
	ID          int    `json:"id"`
	ServiceName string `json:"service"`
	Category_id int    `json:"category_id"`
}

type GetLocations struct {
	ID       int    `json:"id"`
	District string `json:"district"`
	State    string `json:"state"`
}

type BidDetails struct {
	ID          int     `json:"id"`
	Provider    string  `json:"provider"`
	Estimate    float64 `json:"estimate"`
	Description string  `json:"description"`
}
