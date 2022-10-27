package repository

import (
	"context"
	"errors"
	"hendralijaya/user-management-project/helper"
	"hendralijaya/user-management-project/model/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

type UserRepository interface {
	All() []domain.User
	Create(u domain.User) domain.User
	Update(u domain.User, withPassword bool) domain.User
	Delete(u domain.User)
	FindById(id uint) (domain.User, error)
	VerifyCredential(username, email, password string) (domain.User, error)
	FindByEmail(email string) domain.User
	FindByUsername(username string) domain.User
	IsDuplicateEmail(email string) (bool, error)
	CheckThreeAttemptsLogin(username string, email string) bool
}

type UserConnection struct {
	connection *gorm.DB
	mongoDB 	*mongo.Client
}

func NewUserRepository(connection *gorm.DB, mongoDB *mongo.Client) UserRepository {
	return &UserConnection{
		connection: connection,
		mongoDB: mongoDB,
	}
}

func (c *UserConnection) All() []domain.User {
	var users []domain.User
	c.connection.Preload("Role").Find(&users)
	return users
}

func (c *UserConnection) Create(u domain.User) domain.User {
	u.Password = helper.HashAndSalt([]byte(u.Password))
	c.connection.Save(&u)
	return u
}

func (c *UserConnection) Update(u domain.User, withPassword bool) domain.User {
	if withPassword {
		u.Password = helper.HashAndSalt([]byte(u.Password))
		c.connection.Omit("created_at").Save(&u)
	} else {
		c.connection.Omit("created_at").Save(&u)
	}
	return u
}

func (c *UserConnection) Delete(u domain.User) {
	c.connection.Delete(&u)
}

func (c *UserConnection) VerifyCredential(username, email, password string) (domain.User, error) {
	var user domain.User

	if username != "" {
		c.connection.Find(&user, "username = ?", username)
	} else {
		c.connection.Find(&user, "email = ?", email)
	}

	res := helper.ComparedPassword(user.Password, []byte(password))
	if !res {
		failedUser := mongo.IndexModel{
			Keys: bson.M{
				"username": user.Username,
				"password": user.Password,
				"email": user.Email },
			Options: options.Index().SetExpireAfterSeconds(60 * 10),
		}
		c.mongoDB.Database("user-management").Collection("users").Indexes().CreateOne(context.Background(), failedUser)
		return user, errors.New("invalid credential")
	}
	return user, nil
}

func (c *UserConnection) FindByEmail(email string) domain.User {
	var user domain.User
	c.connection.Preload("Role").Find(&user, "email = ?", email)
	return user
}

func (c *UserConnection) FindByUsername(username string) domain.User {
	var user domain.User
	c.connection.Preload("Role").Find(&user, "username = ?", username)
	return user
}

func (c *UserConnection) FindById(id uint) (domain.User, error) {
	var user domain.User
	c.connection.Preload("Role").Find(&user, "id = ?", id)
	if user.ID == 0 {
		return user, errors.New("user id not found")
	}
	return user, nil
}

func (c *UserConnection) IsDuplicateEmail(email string) (bool, error) {
	var user domain.User
	c.connection.Find(&user, "email = ?", email)
	if user.ID == 0 {
		return false, nil
	}
	return true, errors.New("email already exists")
}

func (c *UserConnection) CheckThreeAttemptsLogin(username string, email string) bool {
	userCollection := c.mongoDB.Database("user-management").Collection("users")
	filter := bson.M{
		"username": username,
		"$or": []bson.M{
			{"email": email},
		},
	}
	count, err := userCollection.CountDocuments(context.Background(), filter)
	helper.PanicIfError(err)
	return count >= 3
}
