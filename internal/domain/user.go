package domain

// User represents the aggregate root for user data.
type User struct {
	ID    int64
	Name  string
	Email string
}
