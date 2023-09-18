package domain

type Provider struct {
	ID       int    `json:"id" gorm:"unique;not null"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Phone    string `json:"phone" gorm:"not null"`
	Verified bool   `json:"verified" gorm:"default:false"`
}

type Probook struct {
	ID           int        `json:"id" gorm:"unique;not null"`
	ProID        int        `json:"pro_id"`
	Provider     Provider   `json:"-" gorm:"foreignkey:ProID;constraint:OnDelete:CASCADE"`
	ProfessionID int        `json:"profession_id"`
	Profession   Profession `json:"-" gorm:"foreignkey:ProfessionID;constraint:OnDelete:CASCADE"`
}

type PreferredLocation struct {
	ID         int      `json:"id" gorm:"unique;not null"`
	ProID      int      `json:"pro_id"`
	Provider   Provider `json:"-" gorm:"foreignkey:ProID;constraint:OnDelete:CASCADE"`
	DistrictID int      `json:"district_id"`
	District   District `json:"-" gorm:"foreignkey:DistrictID;constraint:OnDelete:CASCADE"`
}

type ProviderNotification struct {
	ID          int      `json:"id" gorm:"unique;not null"`
	ProID       int      `json:"pro_id"`
	Provider    Provider `json:"-" gorm:"foreignkey:ProID;constraint:OnDelete:CASCADE"`
	Description string   `json:"description"`
	TargetURL   string   `json:"target_url"`
}

type Post struct {
	ID          int      `json:"id" gorm:"unique;not null"`
	ProID       int      `json:"pro_id"`
	Provider    Provider `json:"-" gorm:"foreignkey:ProID;constraint:OnDelete:CASCADE"`
	WorkID      int      `json:"work_id"`
	Work        Work     `json:"-" gorm:"foreignkey:WorkID;constraint:OnDelete:CASCADE"`
	Description string   `json:"description"`
}

type PostImages struct {
	ID     int    `json:"id" gorm:"unique;not null"`
	PostID int    `json:"post_id"`
	Post   Post   `json:"-" gorm:"foreignkey:postID;constraint:OnDelete:CASCADE"`
	Image  string `json:"image"`
}

type IdProof struct {
	ID       int      `json:"id" gorm:"unique;not null"`
	ProID    int      `json:"pro_id"`
	Provider Provider `json:"-" gorm:"foreignkey:ProID;constraint:OnDelete:CASCADE"`
	IdProof  string   `json:"id_proof"`
}
