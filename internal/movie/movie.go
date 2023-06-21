package movie

type Movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	Director string `json:"director"`

	Cast []Actor `json:"cast"`
}

type Actor struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
