package movie

type Movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	Director string `json:"director"`

	Cast []Actor `gorm:"many2many:movie_cast;" json:"cast,omitempty"` //MANY TO MANY RELATIONSHIP
}

type Actor struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (m Movie) String() string {
	return "this is a movie"
}

type Storage interface {
	CreateMovie(movie Movie) (Movie, error)
}

type Service struct {
	Storage Storage
}

func (s *Service) Create(m Movie) (Movie, error) {
	movie, err := s.Storage.CreateMovie(m)

	if err != nil {
		return Movie{}, err
	}

	return movie, nil
}
