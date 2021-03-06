package entities

import (
	"time"
)

type Users struct {
	ID       	uint64
	Username 	string
	Email    	string
	Password 	string
	Name     	string
	Gender   	string
	DOB 		time.Time
	Phone    	string
	UserAddress
}

type UserAddress struct {
	ID             uint64
	UserID         uint64
	BuildingNumber string
	Street         string
	City           string
	Country        string
	PostalCode     string
	Province       string
}

type UserService interface {
	GetUser(username string) (*Users, error)
	Register(payload *Users, street string) (*Users, error)
	Login(payload *Users) (interface{}, error)
	Logout(id, username string) error
	UpdateAccount(payload *Users, username string) (*Users, error)
	UpdateAddress(payload *UserAddress, username string) (*UserAddress, error)
	GetAddress(username string) (*UserAddress, error)
}

type UserMysqlRepositoryInterface interface {
	GetUser(username string) (*Users, error)
	Register(payload *Users, street string) (*Users, error)
	FindByEmail(email string) *Users
	UpdateAccount(payload *Users, id uint64) (*Users, error)
	UpdateAddress(payload *UserAddress, id uint64) (*UserAddress, error)
	GetAddress(id uint64) (*UserAddress, error)
}

type UserScribleRepositoryInterface interface {
	FindUserRefreshToken(userID string) error
	DeleteUserRefreshToken(userID string) error
}