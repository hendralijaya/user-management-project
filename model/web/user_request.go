package web

import "time"

type UserRegisterRequest struct {
	Username    string `json:"username" binding:"required,min=5,max=100"`
	FirstName   string `json:"first_name" binding:"required,min=5,max=100"`
	LastName    string `json:"last_name" binding:"required,min=5,max=100"`
	NIK         string `json:"nik" binding:"required,min=5,max=100,alphanum"`
	Address     string `json:"address" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required,min=11,max=13,numeric"`
	Gender      string `json:"gender" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8"`
	CreatedBy   string `json:"created_by" binding:"required,max=80"`
	RoleId      uint64
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required_without=Username"`
	Username string `json:"username" binding:"required_without=Email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type UserNewPasswordRequest struct {
	Password       string `form:"password" json:"password" binding:"required,min=8"`
	RepeatPassword string `form:"repeat_password" json:"repeat_password" binding:"required,min=8,eqfield=Password"`
}

type UserUpdateRequest struct {
	ID               uint64 `json:"id" binding:"required"`
	RoleId           string `json:"role_id" binding:"required"`
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
