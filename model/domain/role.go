package domain

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(20);not null, unique"`
	Description string `json:"description" gorm:"type:TEXT;not null"`
	CreatedBy   string `json:"created_by" gorm:"type:varchar(80);not null"`
	UpdatedBy   string `json:"updated_by" gorm:"type:varchar(80);not null"`
	DeletedBy   string `json:"deleted_by" gorm:"type:varchar(80);not null"`
}
