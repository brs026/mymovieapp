package service

import "github.com/brs026/mymovieapp/model"

type IMovieService interface {
	GetMovies() ([]model.Movie, error)
	GetMovie(id int) (model.Movie, error)
	DeleteMovie(id int) error
	DeleteMovies() error
	CreateMovie(model.Movie) error
	UpdateMovie(id int, movie model.Movie) error
}
