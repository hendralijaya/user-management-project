package domain

import "time"

type Role struct {
	Id          uint64    `json:"id" gorm:"primary_key:auto_increment"`
	Name        string    `json:"name" gorm:"type:varchar(20);not null, unique"`
	Description string    `json:"description" gorm:"type:varchar(255);not null"`
	Created_by  string    `json:"created_by" gorm:"type:varchar(80);not null"`
	Created_at  time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	Updated_by  string    `json:"updated_by" gorm:"type:varchar(80);not null"`
	Updated_at  time.Time `json:"updated_at" gorm:"type:timestamp;not null"`
}
