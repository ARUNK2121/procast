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
