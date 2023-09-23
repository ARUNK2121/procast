package domain

type Admin struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name" gorm:"validate:required"`
	Email     string `json:"email" gorm:"validate:required"`
	Password  string `json:"password" gorm:"validate:required"`
	Previlege string `json:"previlege" gorm:"previlege:2;default:'normal_admin';previlage IN ('super_admin','normal_admin')"`
}

type Category struct {
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Category  string `json:"category" gorm:"unique;not null"`
	IsDeleted bool   `json:"is_deleted" gorm:"Default:false" `
}

type Profession struct {
	ID         int      `json:"id" gorm:"primaryKey;autoIncrement"`
	Profession string   `json:"profession" gorm:"unique,not null"`
	CategoryID int      `json:"category_id"`
	Category   Category `json:"-" gorm:"foreignkey:CategoryID;constraint:OnDelete:CASCADE"`
	IsDeleted  bool     `json:"is_deleted" gorm:"Default:false" `
}

type State struct {
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	State     string `json:"state" gorm:"unique;not null"`
	IsDeleted bool   `json:"is_deleted" gorm:"Default:false"`
}

type District struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	District string `json:"district" gorm:"unique;not null"`
	StateID  int    `json:"state_id"`
	State    State  `json:"-" gorm:"foreignkey:StateID;constraint:OnDelete:CASCADE"`
}
