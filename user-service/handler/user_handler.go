package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sainath/e-commerce-app/common/pkg"
	"github.com/sainath/e-commerce-app/user-service/model"
	"github.com/sainath/e-commerce-app/user-service/repository"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = repository.DbConnection()
	if err != nil {
		log.Fatalln("Unable to connect database, due to error: " + err.Error())
	}
	db.AutoMigrate(&model.UserAddress{})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var userLogin model.UserLogin
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&userLogin)
	token, err := pkg.GenerateToken(userLogin.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}
