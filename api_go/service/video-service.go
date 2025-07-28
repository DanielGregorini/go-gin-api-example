package service

import (
	"errors"

	"gorm.io/gorm"

	"github.com/DanielGregorini/go-api-gin/entity"
)

type VideoService interface {
	Save(video entity.Video) (entity.Video, error)
	FindAll(page int, pageSize int) (entity.PaginatedResponse[entity.Video], error)
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

func (s *videoService) FindAll(page int, pageSize int) (entity.PaginatedResponse[entity.Video], error) {
	var videos []entity.Video
	var totalItems int64

	if err := s.db.Model(&entity.Video{}).Count(&totalItems).Error; err != nil {
		return entity.PaginatedResponse[entity.Video]{}, err
	}

	offset := (page - 1) * pageSize
	if err := s.db.Limit(pageSize).Offset(offset).Find(&videos).Error; err != nil {
		return entity.PaginatedResponse[entity.Video]{}, err
	}

	totalPages := int((totalItems + int64(pageSize) - 1) / int64(pageSize))

	response := entity.PaginatedResponse[entity.Video]{
		Data:        videos,
		Page:        page,
		PageSize:    pageSize,
		TotalItems:  int(totalItems),
		TotalPages:  totalPages,
		HasNext:     page < totalPages,
		HasPrevious: page > 1,
	}

	return response, nil
}


func (s *videoService) FindByID(id int) (entity.Video, error) {
	var video entity.Video

	if err := s.db.First(&video, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Video{}, errors.New("vídeo não encontrado")
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
