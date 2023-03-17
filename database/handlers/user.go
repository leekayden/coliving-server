package handlers

import (
	"coliving-server/database"
	"coliving-server/database/models"
	"coliving-server/datatypes"
	"fmt"
)

func CreateUser(user *models.User) {
	database.Db.Create(user)
}

func DoesUserExist(id uint64) bool {
	result := datatypes.Msi{}
	database.Db.Table("users").Where("id = ?", id).First(&result)
	fmt.Printf("%#v\n", result)
	return len(result) != 0
}

func DoesEmailExist(email string) (bool, error) {
	result := &models.User{}
	err := database.Db.
		Where("email = ?", email).
		Limit(1).
		Find(&result).
		Error
	if err != nil {
		return true, err
	}
	return result.Id != 0, nil
}


// TODO
// Update user info