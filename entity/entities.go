package entity

import "time"

type Category struct {
	ID          uint
	Name        string
	Description string
	Image       string
}

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
	Name     string
	Message  string
	Email    string
	PlacedAt time.Time
}
