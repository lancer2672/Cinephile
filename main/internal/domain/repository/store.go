package repository

type Store interface {
	MovieRepository

	Disconnect() error
}
