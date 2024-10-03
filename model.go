package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type movie struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Genre  string  `json:"genre"`
	Rating float64 `json:"rating"`
}

func getMovies(db *sql.DB) ([]movie, error) {
	query := "SELECT id, title, genre, rating from movies"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	movies := []movie{}
	for rows.Next() {
		var m movie
		err := rows.Scan(&m.ID, &m.Title, &m.Genre, &m.Rating)
		if err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}

	return movies, nil

}

func (m *movie) getMovie(db *sql.DB) error {
	query := fmt.Sprintf("SELECT title, genre, rating FROM movies WHERE id=%v", m.ID)
	row := db.QueryRow(query)
	err := row.Scan(&m.Title, &m.Genre, &m.Rating)
	if err != nil {
		return err
	}

	return nil
}

func (m *movie) createMovie(db *sql.DB) error {
	query := fmt.Sprintf("INSERT into movies(title, genre, rating) values('%v', '%v', %v)", m.Title, m.Genre, m.Rating)

	result, err := db.Exec(query)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	m.ID = int(id)
	return nil

}

func (m *movie) updateMovie(db *sql.DB) error {
	query := fmt.Sprintf("update movies set title='%v', genre='%v', rating=%v where id=%v", m.Title, m.Genre, m.Rating, m.ID)

	rows, err := db.Exec(query)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("No such rows exists")
	}

	return nil
}

func (m *movie) deleteMovie(db *sql.DB) error {
	query := fmt.Sprintf("DELETE FROM movies where id=%v", m.ID)

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
