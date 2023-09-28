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
