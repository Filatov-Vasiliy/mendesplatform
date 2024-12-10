package storage

import "database/sql"

type Storage struct {
	db *sql.DB
}

func New(url string) (*Storage, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	err = createTables(&Storage{db: db})
	if err != nil {
		return nil, err
	}
	return &Storage{
		db: db,
	}, nil
}
