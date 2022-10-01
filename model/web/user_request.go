package web

import "time"

type UserCreateRequest struct {
	RoleId           uint64 `json:"role_id" binding:"required, numeric"`
	Username         string `json:"username" binding:"required,min=5,max=100"`
	FirstName        string `json:"first_name" binding:"required,min=5,max=100"`
	LastName         string `json:"last_name" binding:"required,min=5,max=100"`
	NIK              string `json:"nik" binding:"required,min=5,max=100,alphanum"`
	Address          string `json:"address" binding:"required"`
	PhoneNumber      string `json:"phone_number" binding:"required, min=16,max=20,numeric"`
	Gender           string `json:"gender" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	Password         string `json:"password" binding:"required,min=8"`
	CreatedBy        string `json:"created_by" binding:"required,max=80"`
	VerificationTime time.Time
}

type UserUpdateRequest struct {
	ID               uint64
	RoleId           uint64 `json:"role_id" binding:"required, numeric"`
	Username         string `json:"username" binding:"required,min=5,max=100"`
	FirstName        string `json:"first_name" binding:"required,min=5,max=100"`
	LastName         string `json:"last_name" binding:"required,min=5,max=100"`
	NIK              string `json:"nik" binding:"required,min=5,max=100,alphanum"`
	Address          string `json:"address" binding:"required"`
	PhoneNumber      string `json:"phone_number" binding:"required, min=16,max=20,numeric"`
	Gender           string `json:"gender" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	Password         string `json:"password" binding:"required,min=8"`
	CreatedBy        string `json:"created_by" binding:"required,max=80"`
	VerificationTime time.Time
}
