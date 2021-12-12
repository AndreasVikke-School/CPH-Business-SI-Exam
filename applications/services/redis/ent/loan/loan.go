// Code generated by entc, DO NOT EDIT.

package loan

import (
	"fmt"
)

const (
	// Label holds the string label denoting the loan type in the database.
	Label = "loan"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEntityId holds the string denoting the entityid field in the database.
	FieldEntityId = "entity_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the loan in the database.
	Table = "loans"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "loans"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_loans"
)

// Columns holds all SQL columns for loan fields.
var Columns = []string{
	FieldID,
	FieldEntityId,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "loans"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_loans",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusRESERVED Status = "RESERVED"
	StatusPICKED   Status = "PICKED"
	StatusPACKED   Status = "PACKED"
	StatusSHIPPED  Status = "SHIPPED"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusRESERVED, StatusPICKED, StatusPACKED, StatusSHIPPED:
		return nil
	default:
		return fmt.Errorf("loan: invalid enum value for status field: %q", s)
	}
}
