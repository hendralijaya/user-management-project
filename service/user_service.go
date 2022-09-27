package service

import (
	"hendralijaya/user-management-project/model/domain"
	"hendralijaya/user-management-project/model/web"
	"hendralijaya/user-management-project/repository"

	"github.com/mashingan/smapping"
)

type UserService interface {
	All() []domain.User
	VerifyCredential(b web.UserLoginRequest) (interface{}, error)
	Create(b web.UserRegisterRequest) (domain.User, error)
	FindById(id uint64) (domain.User, error)
	Update(b web.UserUpdateRequest) (domain.User, error)
	FindByEmail(email string) domain.User
	Delete(id uint64) error
	// Logout(u web.UserLogoutRequest) (domain.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) All() []domain.User {
	return s.userRepository.All()
}

func (s *userService) VerifyCredential(u web.UserLoginRequest) (interface{}, error) {
	user, err := s.userRepository.VerifyCredential(u.Email, u.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) Create(request web.UserRegisterRequest) (domain.User, error) {
	user := domain.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&request))

	if err != nil {
		return user, err
	}

	_, err = s.userRepository.IsDuplicateEmail(request.Email)
	if err != nil {
		return user, err
	}
	return s.userRepository.Create(user), nil
}

func (s *userService) Update(b web.UserUpdateRequest) (domain.User, error) {
	user := domain.User{}
	res, err := s.userRepository.FindById(b.Id)
	if err != nil {
		return user, err
	}
	err = smapping.FillStruct(&user, smapping.MapFields(&b))
	if err != nil {
		return user, err
	}
	user.Email = res.Email
	user.RoleId = res.RoleId
	return s.userRepository.Update(user), nil
}

func (s *userService) FindById(id uint64) (domain.User, error) {
	user, err := s.userRepository.FindById(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) FindByEmail(email string) domain.User {
	user := s.userRepository.FindByEmail(email)
	return user
}

func (s *userService) Delete(id uint64) error {
	user, err := s.userRepository.FindById(id)
	if err != nil {
		return err
	}
	s.userRepository.Delete(user)
	return nil
}
