package service

import (
	"errors"
	"sync"

	"github.com/DanielGregorini/go-api-gin/entity"
)

type VideoService interface {
	Save(video entity.Video) entity.Video
	FindAll() []entity.Video
	FindByID(id int) (entity.Video, error)
	Update(video entity.Video) (entity.Video, error)
	Delete(id int) error
}

type videoService struct {
	mu      sync.Mutex
	videos  []entity.Video
	nextID  int
}

func NewVideoService() VideoService {
	return &videoService{
		videos: make([]entity.Video, 0),
		nextID: 1,
	}
}

func (s *videoService) Save(video entity.Video) entity.Video {
	s.mu.Lock()
	defer s.mu.Unlock()

	// atribui ID incremental
	video.ID = s.nextID
	s.nextID++
	s.videos = append(s.videos, video)
	return video
}

func (s *videoService) FindAll() []entity.Video {
	s.mu.Lock()
	defer s.mu.Unlock()

	// retorna cópia para evitar modificação externa
	result := make([]entity.Video, len(s.videos))
	copy(result, s.videos)
	return result
}

func (s *videoService) FindByID(id int) (entity.Video, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, v := range s.videos {
		if v.ID == id {
			return v, nil
		}
	}
	return entity.Video{}, errors.New("vídeo não encontrado")
}

func (s *videoService) Update(video entity.Video) (entity.Video, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, v := range s.videos {
		if v.ID == video.ID {
			// atualiza campos
			s.videos[i].UserID = video.UserID
			s.videos[i].Title = video.Title
			s.videos[i].Description = video.Description
			s.videos[i].URL = video.URL
			return s.videos[i], nil
		}
	}
	return entity.Video{}, errors.New("vídeo não encontrado")
}

func (s *videoService) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, v := range s.videos {
		if v.ID == id {
			// remove do slice
			s.videos = append(s.videos[:i], s.videos[i+1:]...)
			return nil
		}
	}
	return errors.New("vídeo não encontrado")
}
