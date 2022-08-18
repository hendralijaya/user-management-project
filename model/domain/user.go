package domain

import "time"

type User struct {
	Id               uint64    `json:"id" gorm:"primary_key:auto_increment"`
	Name         string    `json:"username" gorm:"type:varchar(255);not null, unique"`
	Email            string    `json:"email" gorm:"type:varchar(255);not null, unique"`
	Password         string    `json:"-" gorm:"type:varchar(255);not null"`
	Role_id uint64 `json:"role_id" gorm:"type:not null"`
	VerificationTime time.Time `json:"verification_time,omitempty" gorm:"type:timestamp"`
	Token            string    `json:"token,omitempty" gorm:"-"`
}