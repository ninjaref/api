package models

import (
	"database/sql/driver"
	"fmt"
)

// NullableString represents a PostgreSQL text field that can be NULL. Its NULL
// value is the empty string ("").
type NullableString string

// Value returns either the string or "" (if NULL).
func (n NullableString) Value() (driver.Value, error) {
	return string(n), nil
}

// Scan tries to create a NullableString.
func (n *NullableString) Scan(v interface{}) error {
	if v == nil {
		*n = ""
		return nil
	}
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("expected string type, got %T", v)
	}
	*n = NullableString(s)
	return nil
}

// NullableInt represents a PostgreSQL integer field that can be NULL. Its NULL
// value is zero (0).
type NullableInt int64

// Value returns either the integer or 0 (if NULL).
func (n NullableInt) Value() (driver.Value, error) {
	return int64(n), nil
}

// Scan tries to create a NullableInt.
func (n *NullableInt) Scan(v interface{}) error {
	if v == nil {
		*n = 0
		return nil
	}
	s, ok := v.(int64)
	if !ok {
		return fmt.Errorf("expected string type, got %T", v)
	}
	*n = NullableInt(s)
	return nil
}
