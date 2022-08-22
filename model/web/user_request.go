package web

import "time"

type UserRegisterRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=255"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Role_id  uint64
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,min=3,max=255"`
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
	Id               uint64    `json:"id" binding:"required"`
	Name             string    `json:"name" binding:"required,min=3,max=255"`
	Password         string    `json:"password" binding:"required,min=8"`
	RepeatPassword   string    `form:"repeat_password" json:"repeat_password" binding:"required,min=8,eqfield=Password"`
	VerificationTime time.Time `json:"verification_time"`
}
