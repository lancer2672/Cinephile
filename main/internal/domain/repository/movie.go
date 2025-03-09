package repository

import (
	"context"

	"github.com/lancer2672/cinephile/main/internal/domain/model"
)

type MovieRepository interface {
	GetMovies(ctx context.Context, filter model.MovieFilter) ([]model.Movie, int64, error)
	// GetComingSoonMovies(ctx context.Context, limit int) ([]model.Movie, error)

	GetMovieByID(ctx context.Context, id int64) (*model.Movie, error)

	// CreateMovie(ctx context.Context, movie *model.Movie) error

	// UpdateMovie(ctx context.Context, movie *model.Movie) error

	// DeleteMovie(ctx context.Context, id int64) error

	// UpdateMovieStatus(ctx context.Context, id int64, status string) error

	// GetNowShowingMovies(ctx context.Context, limit int) ([]model.Movie, error)

	// UpdateMovieUpdatedAt(ctx context.Context, id int64, updatedAt time.Time) error
}
