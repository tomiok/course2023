package movie

import (
	"fmt"
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	Director string `json:"director"`

	Cast []Actor `gorm:"many2many:movie_cast;" json:"cast,omitempty"` //MANY TO MANY RELATIONSHIP
}

type Actor struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (m Movie) String() string {
	return fmt.Sprintf("Title: %s - Genre: %s - Director: %s", m.Title, m.Genre, m.Director)
}

type Storage interface {
	CreateMovie(movie Movie) (Movie, error)
}

type Service struct {
	Storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		Storage: storage,
	}
}

func (s *Service) Create(m Movie) (Movie, error) {
	movie, err := s.Storage.CreateMovie(m)

	if err != nil {
		return Movie{}, err
	}

	return movie, nil
}
