package entity

import "time"

// User represents application user
type User struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
	Password  string
}

// Comment represents comments forwarded by application users
type Comment struct {
	ID       uint
	UserID   uint
	Name     string
	Message  string
	Email    string
	Rating   uint
	PlacedAt time.Time
}
