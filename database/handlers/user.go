package handlers

import (
	"coliving-server/database"
	"coliving-server/database/models"
	"coliving-server/datatypes"
	"fmt"
	"time"
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

func GetUserById(id uint64) (*models.User, error) {
	user := &models.User{}
	err := database.Db.First(user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(user *models.User) error {
	// Get the location of the database server
	loc, err := time.LoadLocation(database.Db.Config.NowFunc().Location().String())
	if err != nil {
		return err
	}

	err = database.Db.Model(user).Updates(map[string]interface{}{
		"fname":           user.FirstName,
		"lname":           user.LastName,
		"location":        user.Location,
		"email":           user.Email,
		"password":        user.Password,
		"receiveAddEmail": user.ReceiveAddEmails,
		"updated_at":      time.Now().In(loc),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id uint64) error {
	err := database.Db.Delete(&models.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

// TODO
// Update user info
