package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	RoleId           uint64    `json:"role_id" gorm:"type:uint;not null, unsigned"`
	Username         string    `json:"username" gorm:"type:varchar(100);not null"`
	FirstName        string    `json:"first_name" gorm:"type:varchar(100);not null"`
	LastName         string    `json:"last_name" gorm:"type:varchar(100)"`
	NIK              string    `json:"nik" gorm:"type:varchar(100)"`
	Address          string    `json:"address" gorm:"type:text;not null"`
	PhoneNumber      string    `json:"phone_number" gorm:"type:varchar(20);not null"`
	Gender           string    `json:"gender" gorm:"type:varchar(50);not null"`
	Email            string    `json:"email" gorm:"type:varchar(100);not null, unique"`
	Password         string    `json:"-" gorm:"type:varchar(100);not null"`
	CreatedBy        string    `json:"created_by" gorm:"type:varchar(80);not null"`
	VerificationTime time.Time `json:"verification_time,omitempty" gorm:"type:timestamp"`
	AuthToken        string    `json:"auth_token,omitempty" gorm:"-"`
}
