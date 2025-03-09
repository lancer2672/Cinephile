package model

import "time"

type Movie struct {
	ID          int       `json:"id" bson:"id"`
	Title       string    `json:"title" bson:"title"`
	Duration    int       `json:"duration" bson:"duration"` // phút
	AgeLimit    int       `json:"age_limit" bson:"age_limit"`
	Director    string    `json:"director" bson:"director"`
	Actors      []string  `json:"actors" bson:"actors"`
	ReleaseDate time.Time `json:"release_date" bson:"release_date"`
	Description string    `json:"description" bson:"description"`
	Trailer     string    `json:"trailer" bson:"trailer"`
	Country     string    `json:"country" bson:"country"`
	Genre       string    `json:"genre" bson:"genre"`
	Poster      string    `json:"poster" bson:"poster"`
	Status      string    `json:"status" bson:"status"` ///  Status ENUM('Coming Soon', 'Now Showing', 'Ended') DEFAULT 'Coming Soon',
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}
