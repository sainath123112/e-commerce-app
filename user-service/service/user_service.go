package service

import (
	"log"

	"github.com/sainath/e-commerce-app/user-service/model"
	"github.com/sainath/e-commerce-app/user-service/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = repository.DbConnection()
	if err != nil {
		log.Fatalln("Unable to connect databse due to: " + err.Error())
	}

}

func IsUserExist(username string) (bool, error) {
	var user model.User
	err = db.Model(&model.User{}).Where("email = ?", username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false, err
	}
	return true, nil
}

func IsAuthenticated(username string, password string) (bool, error) {
	var user model.User
	db.Model(&model.User{}).Where("email = ?", username).First(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

func RegisterUserService(userRegister model.UserRegisterRequestDto) (bool, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userRegister.Password), 14)
	var user model.User
	if err != nil {
		return false, err
	}
	user.FirstName = userRegister.FirstName
	user.LastName = userRegister.LastName
	user.Email = userRegister.Email
	user.PasswordHash = string(passwordHash)
	err = db.Model(&model.User{}).Create(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
