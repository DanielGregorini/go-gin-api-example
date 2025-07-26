package service

import (
	"errors"
	"gorm.io/gorm"
	"github.com/DanielGregorini/go-api-gin/entity"
)

type VideoService interface {
	Save(video entity.Video) (entity.Video, error)
	FindAll() ([]entity.Video, error)
	FindByID(id int) (entity.Video, error)
	Update(video entity.Video) (entity.Video, error)
	Delete(id int) error
}

type videoService struct {
	db *gorm.DB
}

func NewVideoService(db *gorm.DB) VideoService {
	return &videoService{db: db}
}

func (s *videoService) Save(video entity.Video) (entity.Video, error) {
	if err := s.db.Create(&video).Error; err != nil {
		return entity.Video{}, err
	}
	
	return video, nil
}

func (s *videoService) FindAll() ([]entity.Video, error) {
	var videos []entity.Video

	if err := s.db.Find(&videos).Error; err != nil {
		return nil, err
	}

	return videos, nil
}

func (s *videoService) FindByID(id int) (entity.Video, error) {
	var video entity.Video

	if err := s.db.First(&video, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Video{}, errors.New("vídeo nao encontrado")
		}
		return entity.Video{}, err
	}

	return video, nil
}

func (s *videoService) Update(video entity.Video) (entity.Video, error) {
	if err := s.db.Save(&video).Error; err != nil {
		return entity.Video{}, err
	}

	return video, nil
}

func (s *videoService) Delete(id int) error {
	if err := s.db.Delete(&entity.Video{}, id).Error; err != nil {
		return err
	}

	return nil
}
