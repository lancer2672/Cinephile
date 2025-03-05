package model

import "time"

type Movie struct {
	ID                 int           `json:"id" bson:"id"`
	Name               string        `json:"name" bson:"name"`
	Tags               []string      `json:"tags" bson:"tags"`
	Duration           time.Duration `json:"duration" bson:"duration"`
	AgeLimit           int           `json:"age_limit" bson:"age_limit"`
	Director           string        `json:"director" bson:"director"`
	Actors             []string      `json:"actors" bson:"actors"`
	ReleaseDate        time.Time     `json:"release_date" bson:"release_date"`
	ContentDescription string        `json:"content_description" bson:"content_description"`
	Trailer            string        `json:"trailer" bson:"trailer"`
	Poster             string        `json:"poster" bson:"poster"`
}
