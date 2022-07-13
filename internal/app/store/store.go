package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//DB Config
type Store struct{
	config *Config
	db *sql.DB
}

// New config init
func New(config *Config) *Store{
	return &Store{
		config: config,
	}
}

//open method
func (s *Store) Open() error{
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil{
		return err
	}

	if err := db.Ping(); err != nil{
		return err
	}

	s.db = db
	return nil
}

func (s *Store) Close(){
	s.db.Close()
}