package domain

import "time"

type User struct {
	ID        uint
	Email     string
	Username  string
	FullName  string
	Phone     string
	Address   string
	City      string
	Province  string
	ZipCode   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository interface {
	FindByID(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error
}
