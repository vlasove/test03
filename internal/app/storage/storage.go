package storage

import (
	"database/sql"

	_ "github.com/lib/pq" // postgres adapter
)

// Storage ...
type Storage struct {
	config *Config
	db     *sql.DB // later
	repo   *Repository
}

// New ...
func New(config *Config) *Storage {
	return &Storage{config: config}
}

// Open ...
func (s *Storage) Open() error {
	db, err := sql.Open("postgres", s.config.buildConnStr())
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

// Close ...
func (s *Storage) Close() {
	s.db.Close()
}

// Repo is public manager method for repository
func (s *Storage) Repo() *Repository {
	if s.repo != nil {
		return s.repo
	}
	s.repo = &Repository{storage: s}
	return s.repo
}
