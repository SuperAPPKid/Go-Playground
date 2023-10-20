package user

import (
	"gorm.io/gorm"
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
