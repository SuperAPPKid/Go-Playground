package user

type Gender string

const (
	GenderMale   Gender = "M"
	GenderFemale Gender = "F"
)

type Profile struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(255)"`
	Gender      Gender `gorm:"type:varchar(1)"`
	Email       string `gorm:"type:varchar(255)"`
	PhoneNumber string `gorm:"type:varchar(20)"`

	UserID uint `gorm:"not null"`
}

func (Profile) TableName() string {
	return "Profile"
}
