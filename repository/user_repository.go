package repository

import (
	"errors"
	"hendralijaya/user-management-project/helper"
	"hendralijaya/user-management-project/model/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	All() []domain.User
	Create(u domain.User) domain.User
	Update(u domain.User) domain.User
	Delete(u domain.User)
	FindById(id uint64) (domain.User, error)
	VerifyCredential(email, password string) (domain.User, error)
	FindByEmail(email string) domain.User
	IsDuplicateEmail(email string) (bool, error)
}

type UserConnection struct {
	connection *gorm.DB
}

func NewUserRepository(connection *gorm.DB) UserRepository {
	return &UserConnection{connection: connection}
}

func (c *UserConnection) All() []domain.User {
	var users []domain.User
	c.connection.Find(&users)
	return users
}

func (c *UserConnection) Create(u domain.User) domain.User {
	u.Password = helper.HashAndSalt([]byte(u.Password))
	c.connection.Save(&u)
	return u
}

func (c *UserConnection) Update(u domain.User) domain.User {
	u.Password = helper.HashAndSalt([]byte(u.Password))
	c.connection.Save(&u)
	return u
}

func (c *UserConnection) Delete(u domain.User) {
	c.connection.Delete(&u)
}

func (c *UserConnection) VerifyCredential(email, password string) (domain.User, error) {
	var user domain.User
	c.connection.Find(&user, "email = ?", email)
	res := helper.ComparedPassword(user.Password, []byte(password))
	if !res {
		return user, errors.New("wrong credential")
	}
	return user, nil
}

func (c *UserConnection) FindByEmail(email string) domain.User {
	var user domain.User
	c.connection.Find(&user, "email = ?", email)
	return user
}

func (c *UserConnection) FindById(id uint64) (domain.User, error) {
	var user domain.User
	c.connection.Find(&user, "id = ?", id)
	if user.Id == 0 {
		return user, errors.New("id not found")
	}
	return user, nil
}

func (c *UserConnection) IsDuplicateEmail(email string) (bool, error) {
	var user domain.User
	c.connection.Find(&user, "email = ?", email)
	if user.Id == 0 {
		return false, nil
	}
	return true, errors.New("email already exists")
}
