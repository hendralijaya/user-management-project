package service

import (
	"hendralijaya/user-management-project/model/domain"
	"hendralijaya/user-management-project/model/web"
	"hendralijaya/user-management-project/repository"

	"github.com/mashingan/smapping"
)

type RoleService interface {
	All() []domain.Role
	Create(r web.RoleCreateRequest) (domain.Role, error)
	FindById(id uint) (domain.Role, error)
	Update(r web.RoleUpdateRequest) (domain.Role, error)
	Delete(id uint) error
}

type roleService struct {
	roleRepository repository.RoleRepository
}

func NewRoleService(roleRepository repository.RoleRepository) *roleService {
	return &roleService{roleRepository: roleRepository}
}

func (s *roleService) All() []domain.Role {
	return s.roleRepository.All()
}

func (s *roleService) Create(r web.RoleCreateRequest) (domain.Role, error) {
	role := domain.Role{}
	err := smapping.FillStruct(&role, smapping.MapFields(&r))
	if err != nil {
		return role, err
	}
	return s.roleRepository.Create(role), nil
}

func (s *roleService) FindById(id uint) (domain.Role, error) {
	role, err := s.roleRepository.FindById(uint64(id))
	if err != nil {
		return role, err
	}
	return role, nil
}

func (s *roleService) Update(r web.RoleUpdateRequest) (domain.Role, error) {
	role := domain.Role{}
	_, err := s.roleRepository.FindById(r.ID)
	if err != nil {
		return role, err
	}
	err = smapping.FillStruct(&role, smapping.MapFields(&r))
	if err != nil {
		return role, err
	}
	return s.roleRepository.Update(role), nil
}

func (s *roleService) Delete(id uint) error {
	role, err := s.roleRepository.FindById(uint64(id))
	if err != nil {
		return err
	}
	s.roleRepository.Delete(role)
	return nil
}
