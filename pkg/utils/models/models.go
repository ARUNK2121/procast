package models

import (
	"mime/multipart"

	"github.com/golang-jwt/jwt"
)

type AdminLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ProLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DoubleTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AdminDetailsResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name" `
	Email     string `json:"email" `
	Previlege string `json:"previlege"`
}

type AuthCustomClaims struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Previlege string `json:"previlege"`
	jwt.StandardClaims
}

type VerificationDetails struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	DocumentImage string   `json:"document_images"`
	Services      []string `json:"services"`
}

type Verification struct {
	ID   int
	Name string
}

type ProviderRegister struct {
	Name       string                `json:"name"`
	Email      string                `json:"email"`
	Password   string                `json:"password"`
	RePassword string                `json:"re-password"`
	Phone      string                `json:"phone+6+"`
	Document   *multipart.FileHeader `json:"document"`
}

type UserDetails struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email" `
	Phone     string `json:"phone" `
	IsBlocked bool   `json:"is_blocked"`
}

type ProviderDetails struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	IsVerified bool   `json:"is_verified"`
}

type WorkDetails struct {
	ID            int      `json:"id"`
	Street        string   `json:"street"`
	District      string   `json:"district" `
	State         string   `json:"state" `
	Profession    string   `json:"profession"`
	User          string   `json:"user_name"`
	Provider      string   `json:"provider"`
	Images        []string `json:"images" `
	Participation bool     `json:"participation"`
	WorkStatus    string   `json:"work_status" `
}

type MinWorkDetails struct {
	ID         int    `json:"id"`
	Street     string `json:"street"`
	District   string `json:"district" `
	State      string `json:"state" `
	Profession string `json:"profession"`
	User       string `json:"user_name"`
	Provider   string `json:"provider"`
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

type AddNewState struct {
	State string `json:"state"`
}

type AddNewDistrict struct {
	StateID  int    `json:"state_id"`
	District string `json:"district"`
}

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

type UserSignup struct {
	Name            string `json:"name"`
	Email           string `json:"email" validate:"email"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
}
