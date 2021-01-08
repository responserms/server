// Code generated by entc, DO NOT EDIT.

package token

const (
	// Label holds the string label denoting the token type in the database.
	Label = "token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldExpiredAt holds the string denoting the expired_at field in the database.
	FieldExpiredAt = "expired_at"
	// FieldBlockedAt holds the string denoting the blocked_at field in the database.
	FieldBlockedAt = "blocked_at"

	// EdgeSession holds the string denoting the session edge name in mutations.
	EdgeSession = "session"

	// Table holds the table name of the token in the database.
	Table = "tokens"
	// SessionTable is the table the holds the session relation/edge.
	SessionTable = "sessions"
	// SessionInverseTable is the table name for the Session entity.
	// It exists in this package in order to avoid circular dependency with the "session" package.
	SessionInverseTable = "sessions"
	// SessionColumn is the table column denoting the session relation/edge.
	SessionColumn = "token_session"
)

// Columns holds all SQL columns for token fields.
var Columns = []string{
	FieldID,
	FieldExpiredAt,
	FieldBlockedAt,
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