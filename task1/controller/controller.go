package controller

import "database/sql"

type Controler struct {
	db *sql.DB
}

func Controller(db *sql.DB) *Controler {
	return &Controler{
		db: db,
	}
}
