package storage

import (
	"database/sql"
	"errors"
	"log"

	"github/http/copy/task1/models"
)

func InsertSmedia(db *sql.DB, smedia models.Models) (models.Models, error) {

	query := `
    INSERT INTO smedia (Title, Description, Author_id, Date)
    VALUES ($1, $2, $3, $4)
    `

	_, err := db.Exec(
		query,
		smedia.Title,
		smedia.Description,
		smedia.Author_id,
		smedia.Date,
	)
	if err != nil {
		log.Println("Ошибка при вставке данных1:", err)
		return smedia, err
	}

	return smedia, nil
}

func GetAllSmedia(db *sql.DB) ([]models.Models, error) {
	var smedia []models.Models

	query := `SELECT
		id,
       title,
       description,
       author_id,
       date,
   FROM smedia`

	rows, err := db.Query(query)
	if err != nil {
		log.Println("Ошибка при запросе данных:", err)
		return []models.Models{}, err
	}

	for rows.Next() {
		var sm models.Models
		err = rows.Scan(
			&sm.Title,
			&sm.ID,
			&sm.Description,
			&sm.Author_id,
			&sm.Date,
		)
		if err != nil {
			log.Println("Ошибка при сканировании строки:", err)
			return []models.Models{}, err
		}
		smedia = append(smedia, sm)
	}
	return smedia, nil
}

func GetByIdSmedia(db *sql.DB, id int) (models.Models, error) {
	var smedia models.Models

	query := `SELECT
		id,
		title,
		description,
		author_id,
		date,
	FROM smedia WHERE id=$1`
	err := db.QueryRow(query, id).Scan(
		&smedia.Title,
		&smedia.ID,
		&smedia.Description,
		&smedia.Author_id,
		&smedia.Date,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows", err)
			return smedia, err
		} else {
			log.Println("Ошибка при запросе по ID:", err)
			return smedia, err
		}
	}

	return smedia, nil
}

func UpdateSmedia(db *sql.DB, smedia models.Models) (models.Models, error) {
	query := `UPDATE smedia
	SET title=$3,
		description=$4
		WHERE id=$1 AND author_id=$2`

	_, err := db.Exec(
		query,
		smedia.ID,
		smedia.Author_id,
		smedia.Title,
		smedia.Description,
	)
	if err != nil {
		log.Println("Ошибка при обновлении записи:", err)
		return smedia, err
	}

	return smedia, nil
}

var ErrNotFound = errors.New("smedia not found")

func DeleteSmedia(db *sql.DB, id int) error {
	result, err := db.Exec(`DELETE FROM smedia WHERE id=$1`, id)
	if err != nil {
		log.Println("Ошибка при удалении записи:", err)
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
