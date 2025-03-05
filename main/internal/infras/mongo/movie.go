package mongo

import (
	"context"
	"errors"

	"github.com/lancer2672/cinephile/main/internal/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MovieColleciton = "movies"

// GetMovieByID retrieves a movie by its ID
func (m *MongoStore) GetMovieByID(ctx context.Context, id int) (*model.Movie, error) {
	var movie model.Movie
	opts := options.FindOne().SetProjection(bson.D{{"_id", 0}})
	filter := bson.D{{"id", id}}
	err := m.db.Collection(MovieColleciton).FindOne(ctx, filter, opts).Decode(&movie)
	if err != nil {
		return nil, errors.New("failed to get movie by ID: " + err.Error())
	}
	return &movie, nil
}

// GetMovies retrieves all movies
func (m *MongoStore) GetMovies(ctx context.Context) ([]model.Movie, error) {
	var movies []model.Movie
	opts := options.Find().SetProjection(bson.D{{"_id", 0}})
	cursor, err := m.db.Collection(MovieColleciton).Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, errors.New("failed to get movies: " + err.Error())
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var movie model.Movie
		err := cursor.Decode(&movie)
		if err != nil {
			return nil, errors.New("failed to decode movie: " + err.Error())
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
