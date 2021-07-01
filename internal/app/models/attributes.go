package models

import (
	"database/sql/driver"
	"encoding/json"
)

// Attributes ...
type Attributes struct {
	FirstAttribute  string `json:"first_attribute"`
	SecondAttribute string `json:"second_attribute"`
}

// Value method will call by db.Exec(...)
func (a *Attributes) Value() (driver.Value, error) {
	return json.Marshal(a)
}
