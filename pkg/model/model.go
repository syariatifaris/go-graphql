package model

//Movie structure
type Movie struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Year int64  `json:"year"`
}

//Artist structure
type Artist struct {
	ID   int64  `json:"id"`
	Name string  `json:"name"`
	BOD  string `json:"birth_of_date"`
}

//MovieArtist mapping structure
type MovieArtist struct {
	MovieID  int64 `json:"movie_id"`
	ArtistID int64 `json:"artist_id"`
}
