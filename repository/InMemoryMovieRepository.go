package repository

import (
	"errors"

	"github.com/brs026/mymovieapp/model"
)

type InMemoryMovieRepository struct {
	Movies []model.Movie
}

var (
	ErrMovieNotFound      = errors.New("From repository - movie not found")
	ErrMovieAlreadyExists = errors.New("From repository - movie already exists")
)

func CreateInMemoryMovieRepository() *InMemoryMovieRepository {
	movies := []model.Movie{
		{ID: 1, Title: "The Shawshank Redemption", ReleaseYear: 1994, Score: 9.3},
		{ID: 2, Title: "The Godfather", ReleaseYear: 1972, Score: 9.2},
		{ID: 3, Title: "The Dark Knight", ReleaseYear: 2008, Score: 9.0},
	}

	return &InMemoryMovieRepository{Movies: movies}
}

func (mr *InMemoryMovieRepository) GetMovies() ([]model.Movie, error) {
	if mr == nil {
		return nil, errors.New("Null Pointer")
	}
	return mr.Movies, nil
}

func (mr *InMemoryMovieRepository) GetMovie(id int) (model.Movie, error) {
	for _, movie := range mr.Movies {
		if movie.ID == id {
			return movie, nil
		}
	}
	return model.Movie{}, ErrMovieNotFound
}

func (mr *InMemoryMovieRepository) getMovieIndex(id int) (index int, err error) {
	for index, movie := range mr.Movies {
		if movie.ID == id {
			return index, nil
		}
	}
	return -1, ErrMovieNotFound
}

func (mr *InMemoryMovieRepository) DeleteMovie(id int) (err error) {
	if index, err := mr.getMovieIndex(id); err != nil {
		mr.Movies = append(mr.Movies[:index], mr.Movies[index+1:]...)
	}
	return
}

func (mr *InMemoryMovieRepository) DeleteMovies() error {
	mr.Movies = nil
	return nil
}

func (mr *InMemoryMovieRepository) CreateMovie(movie model.Movie) error {
	movie.ID = len(mr.Movies) + 1
	mr.Movies = append(mr.Movies, movie)
	return nil
}

func (mr *InMemoryMovieRepository) UpdateMovie(id int, movie model.Movie) error {
	if index, err := mr.getMovieIndex(id); err == nil {
		mr.Movies[index] = movie
		return nil
	}
	return ErrMovieNotFound
}
