package service

import (
	"hendralijaya/user-management-project/model/domain"
	"hendralijaya/user-management-project/model/web"
	"hendralijaya/user-management-project/repository"

	"github.com/mashingan/smapping"
)

type UserService interface {
	All() []domain.User
	Create(b web.UserCreateRequest) (domain.User, error)
	FindById(id uint) (domain.User, error)
	Update(b web.UserUpdateRequest) (domain.User, error)
	FindByEmail(email string) domain.User
	Delete(id uint) error
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

func (s *userService) Create(request web.UserCreateRequest) (domain.User, error) {
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
	_, err := s.userRepository.FindById(uint(b.ID))
	if err != nil {
		return user, err
	}
	err = smapping.FillStruct(&user, smapping.MapFields(&b))
	if err != nil {
		return user, err
	}
	if(b.Password != "") {
		return s.userRepository.Update(user,true), nil
	}else {
		return s.userRepository.Update(user,false), nil
	}
}

func (s *userService) FindById(id uint) (domain.User, error) {
	user, err := s.userRepository.FindById(uint(id))
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) FindByEmail(email string) domain.User {
	user := s.userRepository.FindByEmail(email)
	return user
}

func (s *userService) Delete(id uint) error {
	user, err := s.userRepository.FindById(uint(id))
	if err != nil {
		return err
	}
	s.userRepository.Delete(user)
	return nil
}
