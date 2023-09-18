package domain

type Work struct {
	ID               int
	Street           string
	District         string
	State            string
	Country          string
	TargetProfession int
	UserID           int
	ProID            int
	WorkStatus       string
}

type WorkspaceImages struct {
	ID     int
	WorkID int
	Image  int
}

type CompletedImages struct {
	ID     int
	WorkID int
	Image  int
}

type Rating struct {
	ID       int
	Rating   int
	Feedback string
	WorkID   int
}
