package postgres

import "database/sql"

type NewPostgres interface {
	Test() TestRepo
	Login() LoginRepo
}

type store struct {
	testRepo  TestRepo
	loginRepo LoginRepo
	db        *sql.DB
}

func (s *store) Test() TestRepo {
	return s.testRepo
}

func (s *store) Login() LoginRepo {
	return s.loginRepo
}

func NewPostgresI(db *sql.DB) NewPostgres {
	return &store{
		db:        db,
		testRepo:  NewTest(db),
		loginRepo: NewLogin(db),
	}

}
