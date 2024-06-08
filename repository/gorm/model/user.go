package model

import (
	"time"

	"github.com/IbnAnjung/dealls/entity/enuser"
)

type MUser struct {
	ID        string    `gorm:"column:id;primaryKey"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Fullname  string    `gorm:"column:fullname"`
	Gender    uint8     `gorm:"column:gender"`
	Age       uint8     `gorm:"column:age"`
	IsPremium bool      `gorm:"column:is_premium"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (m *MUser) TableName() string {
	return "users"
}

func (m *MUser) ToEntity() (en enuser.User) {
	en.ID = m.ID
	en.Username = m.Username
	en.Fullname = m.Fullname
	en.Password = m.Password
	en.Age = m.Age
	en.Gender = enuser.UserGender(m.Gender)
	en.CreatedAt = m.CreatedAt
	en.IsPremium = m.IsPremium
	return
}

func (m *MUser) FillFromEntity(en enuser.User) {
	m.ID = en.ID
	m.Fullname = en.Fullname
	m.Username = en.Username
	m.Age = en.Age
	m.Password = en.Password
	m.IsPremium = en.IsPremium
	m.Gender = uint8(en.Gender)
	m.CreatedAt = en.CreatedAt
}
