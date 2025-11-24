package storage

import (
	"database/sql"
	"log"

	"github/http/copy/task1/models"
)

func InsertReg(db *sql.DB, reg models.LoginRequest) error {

	hashedPassword, err := HashPassword(reg.Password)
	if err != nil {
		return err
	}

	query := `INSERT INTO authors (author_name, password)
    VALUES ($1, $2)`

	_, err = db.Exec(
		query,
		reg.AuthorName,
		hashedPassword,
	)
	if err != nil {
		return err
	}

	return nil
}

func DeleteLog(db *sql.DB, id int) error {
	result, err := db.Exec(`DELETE FROM authors WHERE id=$1`, id)
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
