package entity

import "time"

// Category represents ordermenu Category
type Category struct {
	ID          uint
	Name        string
	Description string
	Image       string
	//Items       []product
}

// Item represents type of products
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

// Order represents customer order
// type Order struct {
// 	ID       uint
// 	PlacedAt time.Time
// 	UserID   uint
// 	ItemID   uint
// 	Quantity uint
// }

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
