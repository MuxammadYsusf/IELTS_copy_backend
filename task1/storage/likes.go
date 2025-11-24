package storage

import (
	"database/sql"
	"github/http/copy/task1/models"
	"log"
)

func CheckLikeExists(db *sql.DB, authorID int) (bool, bool, error) {
	var currentLiked bool
	query := `SELECT COUNT(*) > 0 FROM likes WHERE author_id = $1`
	err := db.QueryRow(query, authorID).Scan(&currentLiked)
	if err != nil && err != sql.ErrNoRows {
		return false, false, err
	}
	return currentLiked, err == nil, nil
}

func RemoveLike(db *sql.DB, likeReq models.LikeRequest) error {

	query := `
    DELETE FROM likes WHERE id = $1 AND author_id = $2
    `

	_, err := db.Exec(
		query,
		likeReq.ID,
		likeReq.AuthorID,
	)
	if err != nil {
		log.Println("Ошибка при отмене лайка:", err)
		return err
	}

	return nil
}

func AddLike(db *sql.DB, likeReq models.LikeRequest) error {
	query := `
    INSERT INTO likes (id, author_id)
	VALUES ($1, $2)
    `
	_, err := db.Exec(
		query,
		likeReq.ID,
		likeReq.AuthorID,
	)
	if err != nil {
		log.Println("Ошибка при вставке данных1:", err)
		return err
	}

	return nil
}

func CountLikes(db *sql.DB, postID string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM likes WHERE id = $1`
	err := db.QueryRow(query, postID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
