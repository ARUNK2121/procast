package models

import "github.com/golang-jwt/jwt"

type AdminLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Tokens struct {
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
