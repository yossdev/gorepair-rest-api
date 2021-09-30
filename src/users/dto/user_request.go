package dto

import (
	"gorepair-rest-api/src/users/entities"
	"time"
)

type UserRequestLoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequestRegisterBody struct {
	Username string 	`json:"username" validate:"required"`
	Name     string 	`json:"name" validate:"required"`
	Email    string 	`json:"email" validate:"required"`
	Password string 	`json:"password" validate:"required"`
	Gender   string 	`json:"gender" validate:"required"`
	DOB   	 time.Time 	`json:"dob" validate:"required"`
	Phone    string 	`json:"phone" validate:"required"`
	Street   string		`json:"street" validate:"required"`
}

func (req *UserRequestLoginBody) ToDomain() *entities.Users {
	return &entities.Users{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *UserRequestRegisterBody) ToDomain() *entities.Users {
	return &entities.Users{
		Username: req.Username,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
		DOB: 	  req.DOB,
		Gender:   req.Gender,
	}
}