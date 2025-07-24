package service

import (
	"errors"

	"github.com/DanielGregorini/go-api-gin/entity"
	"gorm.io/gorm"
)

type UserService interface {
	Save(user entity.User) (entity.User, error)
	FindAll() ([]entity.User, error)
	FindByID(id int) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	Delete(id int) error
	Login(loginRequest entity.LoginRequest) (entity.User, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func (s *userService) Save(user entity.User) (entity.User, error) {
	if err := s.db.Create(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (s *userService) FindAll() ([]entity.User, error) {
	var users []entity.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) FindByID(id int) (entity.User, error) {
	var user entity.User
	if err := s.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, errors.New("usuário não encontrado")
		}
		return entity.User{}, err
	}
	return user, nil
}

func (s *userService) Update(user entity.User) (entity.User, error) {
	if err := s.db.Save(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (s *userService) Delete(id int) error {
	if err := s.db.Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *userService) Login(req entity.LoginRequest) (entity.User, error) {
	var user entity.User
	if err := s.db.Where("username = ? AND password = ?", req.Username, req.Password).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, errors.New("usuário ou senha incorretos")
		}
		return entity.User{}, err
	}
	return user, nil
}
