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
	Name     string
	Message  string
	Email    string
	PlacedAt time.Time
}
type Product struct {
	ID          uint
	Name        string
	Ptype       string
	Price       float64
	Description string
	CreatedAt   time.Time
	//Categories  Category
	Image string
}
type Balance struct {
	ID          uint
	YourBalance float64
}
