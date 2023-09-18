package database

import (
	"github.com/abrarnaim015/KP-golang-tgs7/config"
	"github.com/abrarnaim015/KP-golang-tgs7/models"
	"github.com/labstack/echo"
)

func GetUsers() (interface{}, error) {
	users := []models.User{}

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func GetUser(id int) (interface{}, error) {
	user := models.User{}

	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(c echo.Context) (interface{}, error) {
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(id int) bool {
	result := config.DB.Where("id = ?", id).Delete(&models.User{})

	if result.Error != nil {
		return false
	}
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func Updateuser(c echo.Context, id int) (interface{}, error) {
	user := models.User{}
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	c.Bind(&user)
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}