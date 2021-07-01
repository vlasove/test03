package storage

import (
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/vlasove/test03/internal/app/models"
)

// Repository ...
type Repository struct {
	storage *Storage
}

// DoSomething ...
func (r *Repository) DoSomething(id int, attr *models.Attributes, timestamp time.Time) error {
	if _, err := r.storage.db.Exec("SELECT test.do_something($1, $2, $3)",
		id,
		attr,
		timestamp,
	); err != nil {
		return err
	}
	return nil
}

// GetDBError ...
func (r *Repository) GetDBError(id int) error {
	if _, err := r.storage.db.Exec("SELECT test.get_db_error($1)", id); err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code >= "50000" {
				return fmt.Errorf(
					"user fail : err code : %s and err msg : %s",
					err.Code,
					err.Message,
				)
			}
			return fmt.Errorf("internal fail : %s", err.Message)

		}
		return err
	}
	return nil
}
