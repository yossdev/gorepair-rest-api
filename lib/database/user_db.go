package database

import (
	"gorepair-rest-api/config"
	"gorepair-rest-api/models"
	"gorepair-rest-api/models/tables"
)

func UserRegister(u models.SignUp) interface{} {
	var user tables.User
	user.Name = u.Name
	user.Email = u.Email
	user.Password = u.Password
	user.Phone = u.Phone
	e := config.DB.Create(&user)
	if e.Error != nil {
		return nil
	}
	return user
}

func UpdateUserAddress(param string, update tables.UserAddress) interface{} {
	var user tables.User
	e := config.DB.First(&user, "id = ?", param)
	if e.Error != nil {
		return nil
	}
	user.Address = update
	config.DB.Save(&user)
	return user
}

func UserLogin(login models.Login) interface{} {
	var user tables.User
	result := config.DB.Where("email = ? AND password = ?", login.Email, login.Password).Preload("Address").Preload("Orders").Preload("Ratings").Find(&user)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil
	}
	return user
}

func GetUsers() (interface{}, error) {
	var users []tables.User
	if e := config.DB.Limit(10).Preload("Address").Preload("Orders").Preload("Ratings").Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func UserDetails(param string) (interface{}, error) {
	var user tables.User
	if e := config.DB.Preload("Address").Preload("Orders").Preload("Ratings").First(&user, "id = ?", param).Error; e != nil {
		return nil, e
	}
	return user, nil
}