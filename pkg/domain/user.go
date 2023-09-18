package domain

type User struct {
	ID       int    `json:"id" gorm:"unique;not null"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"not null"`
	Phone    string `json:"phone" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

type UserNotification struct {
	ID          int    `json:"id" gorm:"unique;not null"`
	UserID      int    `json:"user_id"`
	User        User   `json:"-" gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE"`
	Description string `json:"description"`
	TargetURL   string `json:"target_url"`
}
