package enauth

import "time"

type LoginInput struct {
	Username string
	Password string
}

type LoginOutput struct {
	ID           string
	Username     string
	Fullname     string
	Gender       int8
	Age          uint8
	IsPremium    bool
	CreatedAt    time.Time
	AccessToken  string
	RefreshToken string
}
