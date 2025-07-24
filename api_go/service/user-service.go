package service

import (
	"github.com/DanielGregorini/go-api-gin/entity"
	"errors"
)

type UserService interface {
	Save(entity.User) entity.User
	FindAll() []entity.User
	Login(loginRequest entity.LoginRequest) (entity.User, error)
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

func (service *userService) Login(loginRequest entity.LoginRequest) (entity.User, error) {
	for _, user := range service.users {
		if user.Username == loginRequest.Username && user.Password == loginRequest.Password {
			return user, nil
		}
	}

	 return entity.User{}, errors.New("usuário ou senha incorretos")
}