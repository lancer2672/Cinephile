package repository

import (
	"context"

	"github.com/lancer2672/cinephile/main/internal/domain/model"
)

type MovieRepository interface {
	GetMovies(context.Context) ([]model.Movie, error)
	GetMovieByID(context.Context, int) (*model.Movie, error)
}
