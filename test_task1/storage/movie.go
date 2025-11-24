package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github/muhammadyusuf/http/models"

	"github.com/google/uuid"
)

func InsertMovie(db *sql.DB, movie models.Movie) (string, error) {
	var id = uuid.New().String()

	movie.ID = id

	query := `INSERT INTO movie(
	id,         
	title,      
	"duration",    
	description)
	Values($1, $2, $3, $4) `

	_, err := db.Exec(
		query,
		id,
		movie.Title,
		movie.Duration,
		movie.Description,
	)
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}

	return movie.ID, nil
}

func GetByIdMovie(db *sql.DB, id string) (models.Movie, error) {
	var movie models.Movie
	query := `SELECT 
		id,
		title,
		duration,
		description 
		FROM movie WHERE id = $1; `

	log.Printf("Executing query: %s with id: %s", query, id)

	err := db.QueryRow(query, id).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Duration,
		&movie.Description,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows found for id", id)
		} else {
			log.Println("ERROR is here", err, id)
		}
		return models.Movie{}, err
	}

	return movie, nil
}

func GetAllMovie(db *sql.DB) ([]models.Movie, error) {

	var movie []models.Movie

	query := `SELECT 
		id,
		title,
		duration,
		description 
		FROM movie`

	rows, err := db.Query(query)
	if err != nil {
		return []models.Movie{}, err
	}

	for rows.Next() {
		var m models.Movie
		err := rows.Scan(
			&m.ID,
			&m.Title,
			&m.Duration,
			&m.Description,
		)
		if err != nil {
			return []models.Movie{}, err
		}

		movie = append(movie, m)
	}

	return movie, nil
}

func UpdateMovie(db *sql.DB, movie models.Movie) error {

	query := `UPDATE movie
	SET
		title=$2,
		duration=$3,
		description=$4 
	WHERE id = $1`

	_, err := db.Exec(
		query,
		movie.ID,
		movie.Title,
		movie.Duration,
		movie.Description,
	)
	if err != nil {
		return err
	}

	return nil
}

var ErrNotFound = errors.New("movie not found")

func DeleteMovie(db *sql.DB, id string) error {

	result, err := db.Exec("DELETE FROM movie WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
