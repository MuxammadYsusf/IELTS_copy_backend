package storage

import "database/sql"

type NewStorage interface {
	Question() QuestionRepo
	Login() LoginRepo
}

type store struct {
	questionRepo QuestionRepo
	loginRepo    LoginRepo
	db           *sql.DB
}

func (s *store) Question() QuestionRepo {
	return s.questionRepo
}

func (s *store) Login() LoginRepo {
	return s.loginRepo
}
func StorageI(db *sql.DB) NewStorage {
	return &store{
		db:           db,
		questionRepo: NewQuestion(db),
		loginRepo:    NewSession(db),
	}
}
