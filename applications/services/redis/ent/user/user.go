// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// EdgeLoans holds the string denoting the loans edge name in mutations.
	EdgeLoans = "loans"
	// Table holds the table name of the user in the database.
	Table = "users"
	// LoansTable is the table that holds the loans relation/edge.
	LoansTable = "loans"
	// LoansInverseTable is the table name for the Loan entity.
	// It exists in this package in order to avoid circular dependency with the "loan" package.
	LoansInverseTable = "loans"
	// LoansColumn is the table column denoting the loans relation/edge.
	LoansColumn = "user_loans"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldAge,
	FieldPassword,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// AgeValidator is a validator for the "age" field. It is called by the builders before save.
	AgeValidator func(int64) error
)
