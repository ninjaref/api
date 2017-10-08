package models

import (
	"github.com/go-gorp/gorp"
)

// Ninja represents an individual ANW competitor.
type Ninja struct {
	ID         int64          `db:"ninja_id, primarykey, autoincrement"`
	FirstName  string         `db:"first_name"`
	LastName   string         `db:"last_name"`
	Sex        string         `db:"sex"`
	Age        NullableInt    `db:"age"`
	Occupation NullableString `db:"occupation"`
	Instagram  NullableString `db:"instagram"`
	Twitter    NullableString `db:"twitter"`
}

// All returns all competitors in the database.
func (n *Ninja) All(db *gorp.DbMap) (ninjas []Ninja, err error) {
	_, err = db.Select(&ninjas, "SELECT * FROM Ninja")
	return ninjas, err
}
