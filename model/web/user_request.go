package web

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
	Password       string `json:"password" binding:"required,min=8"`
	RepeatPassword string `json:"repeat_password" binding:"required,min=8,eqfield=password"`
}
