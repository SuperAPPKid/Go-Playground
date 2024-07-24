package user

import (
	"github.com/SuperAPPKid/Go-Playground/REST-API/service/postgresql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	Profile   Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt int64   `gorm:"autoCreateTime"`
	gorm.DeletedAt
}

func (User) TableName() string {
	return "User"
}

func FetchAll() ([]*User, error) {
	var users []*User
	r := postgresql.Start().Preload(clause.Associations).Find(&users)
	return users, r.Error
}

func FetchByID(id int) (*User, error) {
	user := &User{ID: uint(id)}
	r := postgresql.Start().Preload(clause.Associations).First(user)
	return user, r.Error
}

func (u *User) Create() error {
	r := postgresql.Start().Create(u)
	return r.Error
}

func (u *User) Delete() error {
	r := postgresql.Start().Delete(u)
	return r.Error
}
