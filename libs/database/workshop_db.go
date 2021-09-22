package database

import (
	"gorepair-rest-api/config"
	"gorepair-rest-api/models"
	"gorepair-rest-api/models/tables"
)

func WorkshopRegister(u models.SignUp) interface{} {
	var workshop tables.Workshop
	workshop.Name = u.Name
	workshop.Email = u.Email
	workshop.Password = u.Password
	workshop.Phone = u.Phone
	e := config.DB.Create(&workshop)
	if e.Error != nil {
		return nil
	}
	return workshop
}

func UpdateWorkshopAddress(param string, update tables.WorkshopAddress) interface{} {
	var workshop tables.Workshop
	e := config.DB.First(&workshop, "id = ?", param)
	if e.Error != nil {
		return nil
	}
	workshop.Address = update
	config.DB.Save(&workshop)
	return workshop
}

func WorkshopLogin(login models.Login) interface{} {
	var workshop tables.Workshop
	result := config.DB.Where("email = ? AND password = ?", login.Email, login.Password).Preload("Address").Preload("Orders").Preload("Ratings").Find(&workshop)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil
	}
	return workshop
}

func GetWorkshops() (interface{}, error) {
	var workshop []tables.Workshop
	if e := config.DB.Limit(10).Preload("Address").Preload("Orders").Preload("Ratings").Find(&workshop).Error; e != nil {
		return nil, e
	}
	return workshop, nil
}

func WorkshopDetails(param string) (interface{}, error) {
	var workshop tables.Workshop
	if e := config.DB.Preload("Address").Preload("Orders").Preload("Ratings").First(&workshop, "id = ?", param).Error; e != nil {
		return nil, e
	}
	return workshop, nil
}

func FindWorkshop(param string) interface{} {
	var workshop []tables.Workshop
	e := config.DB.Limit(10).Where("name LIKE ?", "%"+param+"%").Preload("Address").Preload("Services").Preload("Orders").Preload("Ratings").Find(&workshop)
	if e.Error != nil {
		return nil
	}
	return workshop
}

func UpdateWorkshopDescription(param string, update tables.Description) interface{} {
	var workshop tables.Workshop
	e := config.DB.First(&workshop, "id = ?", param)
	if e.Error != nil {
		return nil
	}
	workshop.Description = update
	config.DB.Save(&workshop)
	return workshop
}