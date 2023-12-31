package domain

type Work struct {
	ID                 int        `json:"id" gorm:"primaryKey;autoIncrement"`
	Street             string     `json:"street"`
	DistrictID         int        `json:"district_id"`
	District           District   `json:"-" gorm:"foreignkey:DistrictID;constraint:OnDelete:CASCADE"`
	StateID            int        `json:"state_id"`
	State              State      `json:"-" gorm:"foreignkey:StateID;constraint:OnDelete:CASCADE"`
	TargetProfessionID int        `json:"target_profession"`
	Profession         Profession `json:"-" gorm:"foreignkey:TargetProfessionID;constraint:OnDelete:CASCADE"`
	UserID             int        `json:"user_id"`
	User               User       `json:"-" gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE"`
	ProID              int        `json:"pro_id"`
	Provider           Provider   `json:"-" gorm:"foreignkey:ProID;constraint:OnDelete:CASCADE"`
	WorkStatus         string     `json:"work_status" gorm:"column:work_status;default:'listed';check:work_status IN ('listed','committed','completed')"`
}

type WorkspaceImages struct {
	ID     int    `json:"id" gorm:"primaryKey;autoIncrement"`
	WorkID int    `json:"work_id"`
	Work   Work   `json:"-" gorm:"foreignkey:WorkID;constraint:OnDelete:CASCADE"`
	Image  string `json:"image"`
}

type CompletedImages struct {
	ID     int    `json:"id" gorm:"primaryKey;autoIncrement"`
	WorkID int    `json:"work_id"`
	Work   Work   `json:"-" gorm:"foreignkey:WorkID;constraint:OnDelete:CASCADE"`
	Image  string `json:"image"`
}

type Rating struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Rating   int    `json:"rating" gorm:"rating"`
	Feedback string `json:"feedback"`
	WorkID   int    `json:"work_id"`
	Work     Work   `json:"-" gorm:"foreignkey:WorkID;constraint:OnDelete:CASCADE"`
}

type Bid struct {
	ID          int      `json:"id" gorm:"primaryKey;autoIncrement"`
	WorkID      int      `json:"work_id"`
	Work        Work     `json:"-" gorm:"foreignkey:WorkID;constraint:OnDelete:CASCADE"`
	ProID       int      `json:"pro_id"`
	Provider    Provider `json:"-" gorm:"foreignkey:ProID;constraint:OnDelete:CASCADE"`
	Estimate    float64  `json:"estimate"`
	Description string   `json:"description"`
}
