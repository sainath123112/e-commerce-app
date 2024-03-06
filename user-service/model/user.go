package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName    string      `gorm:"not null" json:"first_name"`
	LastName     string      `gorm:"not null" json:"last_name"`
	Email        string      `gorm:"primarykey" json:"email"`
	PasswordHash string      `gorm:"not null" json:"password_hash"`
	UserDetails  UserAddress `json:"address"`
}

type UserAddress struct {
	gorm.Model
	UserID     uint   `gorm:"not null" json:"user_id"`
	Address1   string `gorm:"not null" json:"address_1"`
	Address2   string `json:"address_2"`
	City       string `gorm:"not null" json:"city"`
	State      string `gorm:"not null" json:"state"`
	Country    string `gorm:"not null" json:"country"`
	PostalCode string `gorm:"not null" json:"postal_code"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
