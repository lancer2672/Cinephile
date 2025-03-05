package app

import (
	"context"

	"github.com/lancer2672/cinephile/main/internal/domain/model"
	"github.com/lancer2672/cinephile/main/internal/domain/repository"
)

type MovieService interface {
	GetMovies(context.Context) ([]model.Movie, error)
	GetMovieByID(context.Context, int) (*model.Movie, error)
}

type movieService struct {
	store repository.Store
}

func NewMovieApp(store repository.Store) MovieService {
	return &movieService{store: store}
}

func (m *movieService) GetMovies(ctx context.Context) ([]model.Movie, error) {
	return m.store.GetMovies(context.TODO())
}

func (m *movieService) GetMovieByID(ctx context.Context, id int) (*model.Movie, error) {
	return m.store.GetMovieByID(context.TODO(), id)
}
