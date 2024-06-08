package enauth

import "time"

type RegisterInput struct {
	Username        string
	Password        string
	ConfirmPassword string
	Fullname        string
	Gender          *uint8
	Age             uint8
}

type RegisterOutput struct {
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
