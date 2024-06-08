package enuser

import (
	"time"
)

type UserGender uint8
type SwipeType uint8

const (
	UserGenderMale   UserGender = 1
	UserGenderFemale UserGender = 0
	SwipeTypeLike    SwipeType  = 1
	SwipeTypeDislike SwipeType  = 0
)

type User struct {
	ID        string
	Username  string
	Password  string
	Fullname  string
	Gender    UserGender
	Age       uint8
	IsPremium bool
	CreatedAt time.Time
}

func (g *UserGender) GetOposite() UserGender {
	if *g == UserGenderFemale {
		return UserGenderMale
	}

	return UserGenderFemale
}
