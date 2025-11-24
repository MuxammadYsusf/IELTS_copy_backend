package models

import "time"

type Models struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author_id   int    `json:"author_id"`
	Date        int64  `json:"date"`
}

type LikeRequest struct {
	AuthorID int `json:"author_id"`
	ID       int `json:"id"`
}

type LoginRequest struct {
	AuthorID   int    `json:"author_id"`
	AuthorName string `json:"author_name"`
	Password   string `json:"password"`
}

func NewModels(title, description string, author_id, id int, date time.Time, like int) Models {
	return Models{
		Title:       title,
		Description: description,
		Author_id:   author_id,
		ID:          id,
		Date:        date.UnixMilli(),
	}

}
