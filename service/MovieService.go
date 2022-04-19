package service

import (
	"errors"

	"github.com/brs026/mymovieapp/model"
	"github.com/brs026/mymovieapp/repository"
)

var (
	ErrIdNotValid    = errors.New("ID Value not valid")
	ErrMovieNotFound = errors.New("The movie cannot be found")
	ErrEmptyTitle    = errors.New("Title is empty")
)

type MovieService struct {
	movieRepository repository.IMovieRepository
}

func CreateMovieService(movieRepository repository.IMovieRepository) *MovieService {
	return &MovieService{movieRepository: movieRepository}
}

func (ms *MovieService) GetMovies() ([]model.Movie, error) {
	return ms.movieRepository.GetMovies()
}

func (ms *MovieService) GetMovie(id int) (model.Movie, error) {
	if id < 0 {
		return model.Movie{}, ErrIdNotValid
	}
	movie, err := ms.movieRepository.GetMovie(id)

	if err != nil {
		if errors.Is(err, repository.ErrMovieNotFound) {
			return model.Movie{}, ErrMovieNotFound
		}
		return model.Movie{}, err
	}

	return movie, nil
}

func (ms *MovieService) DeleteMovie(id int) error {
	if id < 0 {
		return ErrIdNotValid
	}

	if err := ms.DeleteMovie(id); err != nil {
		if errors.Is(err, repository.ErrMovieNotFound) {
			return ErrMovieNotFound
		}
		return err
	}
	return nil
}

func (ms *MovieService) DeleteMovies() error {
	return ms.movieRepository.DeleteMovies()
}

func (ms *MovieService) CreateMovie(movie model.Movie) error {
	if movie.Title == "" {
		return ErrEmptyTitle
	}
	return ms.movieRepository.CreateMovie(movie)
}

func (ms *MovieService) UpdateMovie(id int, movie model.Movie) error {
	if id < 0 {
		return ErrIdNotValid
	}
	if movie.Title == "" {
		return ErrEmptyTitle
	}
	err := ms.movieRepository.UpdateMovie(id, movie)
	if errors.Is(err, repository.ErrMovieNotFound) {
		return ErrMovieNotFound
	}
	return nil
}
