package user

import (
	"errors"
	"github.com/SuperAPPKid/Go-Playground/REST-API/service/postgresql"
)

type Gender string

const (
	GenderMale   Gender = "M"
	GenderFemale Gender = "F"
)

func (g Gender) Verify() error {
	switch g {
	case GenderMale, GenderFemale:
		return nil
	default:
		return errors.New("Invalid Gender")
	}
}

type Profile struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(255);not null"`
	Gender      Gender `gorm:"type:varchar(1);not null"`
	Email       string `gorm:"type:varchar(255);not null"`
	PhoneNumber string `gorm:"type:varchar(20);not null"`

	UserID uint `gorm:"not null"`
}

func (Profile) TableName() string {
	return "Profile"
}

func (p *Profile) UpdateAll() error {
	r := postgresql.Start().
		Model(&Profile{}).
		Where("user_id = ?", p.UserID).
		Updates(map[string]interface{}{
			"name":         p.Name,
			"gender":       p.Gender,
			"email":        p.Email,
			"phone_number": p.PhoneNumber,
		})
	return r.Error
}

func (p *Profile) UpdateNonZero() error {
	r := postgresql.Start().
		Model(&Profile{}).
		Where("user_id = ?", p.UserID).
		Updates(p)
	return r.Error
}
