package app

import (
	"context"

	"github.com/lancer2672/cinephile/main/internal/domain/model"
	"github.com/lancer2672/cinephile/main/internal/domain/repository"
)

type MovieService interface {
	GetMovies(context.Context, model.MovieFilter) ([]model.Movie, int64, error)
	GetMovieByID(context.Context, int64) (*model.Movie, error)
}

type movieService struct {
	store repository.Store
}

func NewMovieApp(store repository.Store) MovieService {
	return &movieService{store: store}
}

func (m *movieService) GetMovies(ctx context.Context, filter model.MovieFilter) ([]model.Movie, int64, error) {

	movies, totalItem, err := m.store.GetMovies(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	_ = totalItem
	return movies, 0, nil
}

func (m *movieService) GetMovieByID(ctx context.Context, id int64) (*model.Movie, error) {
	// return m.store.GetMovieByID(context.TODO(), id)
	return nil, nil
}
