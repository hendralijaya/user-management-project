package domain

import "time"

type User struct {
	Id               uint64    `json:"id" gorm:"primary_key:auto_increment"`
	RoleId           uint64    `json:"role_id" gorm:"type:uint;not null, unsigned"`
	Name             string    `json:"name" gorm:"type:varchar(255);not null, unique"`
	Email            string    `json:"email" gorm:"type:varchar(255);not null, unique"`
	Password         string    `json:"-" gorm:"type:varchar(100);not null"`
	PhoneNumber      string    `json:"phone_number" gorm:"type:varchar(15);not null"`
	Address          string    `json:"address" gorm:"type:text;not null"`
	Gender           string    `json:"gender" gorm:"type:varchar(50);not null"`
	BirthPlace       string    `json:"birth_place" gorm:"type:varchar(150);not null"`
	BirthDate        time.Time `json:"birth_date" gorm:"type:date;not null"`
	BloodType        string    `json:"blood_type" gorm:"type:varchar(2);not null"`
	KtpNumber        string    `json:"ktp_number" gorm:"type:varchar(20);not null"`
	EmployeeType     string    `json:"employee_type" gorm:"type:varchar(255);not null"`
	CreatedAt        time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	CreatedBy        string    `json:"created_by" gorm:"type:varchar(80);not null"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"type:timestamp;not null"`
	UpdatedBy        string    `json:"updated_by" gorm:"type:varchar(80);not null"`
	StatusResource   string    `json:"status_resource" gorm:"type:varchar(50);not null"`
	VerificationTime time.Time `json:"verification_time,omitempty" gorm:"type:timestamp"`
	AuthToken        string    `json:"auth_token,omitempty" gorm:"-"`
}
