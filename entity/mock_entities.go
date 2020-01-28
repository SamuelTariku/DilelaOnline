package entity

import (
	"time"
)

var ProductMock = Product{
	ID:          1,
	Name:        "Mock Product 01",
	Ptype:       "Mock Product 01 type",
	Price:       30.4,
	Description: "Mock Product 01 description",
	CreatedAt:   time.Time{},

	Image:  "Mock Product 01 Image",
	UserID: 32,
}

var BalanceMock = Balance{
	ID:          1,
	YourBalance: 30002.3,
}

var OrderMock = Order{
	ID:       1,
	PlacedAt: time.Time{},
	UserID:   12,
	ItemID:   23,
}

// User represents application user
var UserMock = User{
	ID:        1,
	FirstName: "Mock User 01",
	LastName:  "Mock User 02",
	Email:     "mockuser@gmail.com",
	Password:  "P@$$w0rd",
}

// Comment represents comments forwarded by application users
var CommentMock = Comment{
	ID:        1,
	Name:      "Mock Comment 01",
	Message:   "mock message 01",
	Email:     "mockuser@gmail.com",
	PlacedAt:  time.Time{},
	UserID:    3,
	ProductID: 34,
	Rating:    4,
}

var AdMock = Advertisement{
	ID:        1,
	ProductID: 23,
	Ptype:      "mock advert 01",

}
//
var CartMock = Cart{
	ID:       1,
	ProductID: 23,
	UserID:   21,
	AddedTime: time.Time{},
	Price:     323,

}

//var SessionMock = Session {
//	ID:         1,
//	UUID:       "_session_one",
//	Expires:    0,
//	Signingkey: []byte("DilelaOnlineApp"),
//}
