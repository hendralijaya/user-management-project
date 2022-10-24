package service

import (
	"errors"
	"hendralijaya/user-management-project/model/domain"
	"hendralijaya/user-management-project/model/web"
	"hendralijaya/user-management-project/repository"

	"github.com/mashingan/smapping"
)

type AuthService interface {
	Login(b web.UserLoginRequest) (domain.User, error)
	Register(b web.UserRegisterRequest) (domain.User, error)
	VerifyRegisterToken(b web.UserRegisterVerificationTokenRequest) (domain.User, error)
	VerifyForgotPasswordToken(b web.UserNewPasswordRequest) (domain.User, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{userRepository: userRepository}
}

func (service *authService) Login(u web.UserLoginRequest) (domain.User, error) {
	threeAttempt := service.userRepository.CheckThreeAttemptsLogin(u.Username, u.Email)
	if threeAttempt {
		return domain.User{}, errors.New("You have reached the maximum number of login attempts. Please try again later")
	}
	user, err := service.userRepository.VerifyCredential(u.Username, u.Email, u.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (service *authService) Register(request web.UserRegisterRequest) (domain.User, error) {
	user := domain.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&request))

	if err != nil {
		return user, err
	}

	_, err = service.userRepository.IsDuplicateEmail(request.Email)
	if err != nil {
		return user, err
	}
	return service.userRepository.Create(user), nil
}

func (service *authService) VerifyRegisterToken(request web.UserRegisterVerificationTokenRequest) (domain.User, error) {
	user := domain.User{}
	user, err := service.userRepository.FindById(uint(request.ID))
	if err != nil {
		return user, err
	}
	user.VerificationTime = request.VerificationTime
	return service.userRepository.Update(user, false), nil
}

func (service *authService) VerifyForgotPasswordToken(request web.UserNewPasswordRequest) (domain.User, error) {
	user := domain.User{}
	user, err := service.userRepository.FindById(uint(request.ID))
	if err != nil {
		return user, err
	}
	user.Password = request.Password
	return service.userRepository.Update(user, true), nil
}
