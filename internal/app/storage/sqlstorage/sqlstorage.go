package sqlstorage

import "database/sql"

type Storage struct {
	db       *sql.DB
	userRepo *UserRepo
}

func (s *Storage) Users() {}
