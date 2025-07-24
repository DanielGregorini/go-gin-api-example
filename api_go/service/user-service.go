package service

import (
	"github.com/DanielGregorini/go-api-gin/entity"
	"errors"
)

type UserService interface {
	Save(entity.User) entity.User
	FindAll() []entity.User
	Login(username, password string) (entity.User, error)
}

type userService struct {
	users []entity.User
}

func NewUserService() UserService {
	return &userService{}
}

func (service *userService) Save(user entity.User) entity.User {
	service.users = append(service.users, user)
	return user
}

func (service *userService) FindAll() []entity.User {
	return service.users
}

func (service *userService) Login(username, password string) (entity.User, error) {
	for _, user := range service.users {
		if user.Username == username && user.Password == password {
			return user, nil
		}
	}

	 return entity.User{}, errors.New("usuário ou senha incorretos")
}