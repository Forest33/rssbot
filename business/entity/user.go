package entity

import "time"

// User is a bot user
type User struct {
	ID        int64
	FirstName string
	LastName  string
	UserName  string
	Language  string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

// UsersFilter is filter for getting users
type UsersFilter struct {
	ID []int64
}
