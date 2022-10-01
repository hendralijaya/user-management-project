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
	ID          uint64
	Username    string
	FirstName   string
	LastName    string
	NIK         string
	Address     string
	PhoneNumber string
	Gender      string
	Email       string `json:"email" binding:"required,email"`
	Password    string
	CreatedBy   string
	RoleId      uint64
}

type UserNewPasswordRequest struct {
	ID             uint64
	Username       string
	FirstName      string
	LastName       string
	NIK            string
	Address        string
	PhoneNumber    string
	Gender         string
	Email          string
	CreatedBy      string
	RoleId         uint64
	Password       string `form:"password" json:"password" binding:"required,min=8"`
	RepeatPassword string `form:"repeat_password" json:"repeat_password" binding:"required,min=8,eqfield=Password"`
}

type UserRegisterVerificationTokenRequest struct {
	ID               uint64
	RoleId           uint64
	Username         string
	FirstName        string
	LastName         string
	NIK              string
	Address          string
	PhoneNumber      string
	Gender           string
	Email            string
	Password         string
	CreatedBy        string
	VerificationTime time.Time
}
