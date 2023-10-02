package domain

import "time"

type User struct {
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name" gorm:"not null"`
	Email     string `json:"email" gorm:"not null"`
	Phone     string `json:"phone" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	IsBlocked bool   `json:"is_blocked" gorm:"Default:false"`
}

type UserNotification struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      int       `json:"user_id"`
	User        User      `json:"-" gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE"`
	Time        time.Time `json:"time"`
	Description string    `json:"description"`
	TargetURL   string    `json:"target_url"`
	IsRead      bool      `json:"is_read" gorm:"default:false"`
}
