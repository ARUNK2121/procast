package models

type UserDetails struct {
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name" gorm:"not null"`
	Email     string `json:"email" gorm:"not null"`
	Phone     string `json:"phone" gorm:"unique;not null"`
	IsBlocked bool   `json:"is_blocked" gorm:"Default:false"`
}
