package models

type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Duration    string `json:"duration"`
	Description string `json:"description"`
}
