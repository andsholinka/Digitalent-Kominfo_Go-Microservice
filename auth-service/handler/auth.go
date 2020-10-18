package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gorm.io/gorm"

	"github.com/andsholinka/Digitalent-Kominfo_Go-Microservice/auth-service/database"
	// "github.com/andsholinka/Digitalent-Kominfo_Go-Microservice/auth-service/utils"
	"github.com/andsholinka/Digitalent-Kominfo_Go-Microservice/utils"
)

type AuthDB struct {
	Db *gorm.DB
}

func (db *AuthDB) ValidateAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	authToken := r.Header.Get("Authorization")

	res, err := database.ValidateAuth(authToken, db.Db)
	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	utils.WrapAPIData(w, r, database.Auth{
		Username: res.Username,
		Token:    res.Token,
	}, http.StatusOK, "success")

	// if authToken == "" {
	// 	utils.WrapAPIError(w, r, "Invalid auth", http.StatusForbidden)
	// 	return
	// }

	// if authToken != "respecker" {
	// 	utils.WrapAPIError(w, r, "Invalid auth", http.StatusForbidden)
	// 	return
	// }

	// utils.WrapAPISuccess(w, r, "success", 200)
}

func (db *AuthDB) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	//TODO Buat Signup

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		utils.WrapAPIError(w, r, "cannot read body", http.StatusBadRequest)
		return
	}

	var signup database.Auth

	err = json.Unmarshal(body, &signup)
	if err != nil {
		utils.WrapAPIError(w, r, "error unmarshal : "+err.Error(), http.StatusInternalServerError)
		return
	}

	signup.Token = utils.IdGenerator()
	err = signup.SignUp(db.Db)
	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	utils.WrapAPISuccess(w, r, "Success", http.StatusOK)
	return

}

func (db *AuthDB) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		utils.WrapAPIError(w, r, "cannot read body", http.StatusBadRequest)
		return
	}

	var login database.Auth

	err = json.Unmarshal(body, &login)
	if err != nil {
		utils.WrapAPIError(w, r, "error unmarshal : "+err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := login.Login(db.Db)
	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	utils.WrapAPIData(w, r, database.Auth{
		Username: res.Username,
		Token:    res.Token,
	}, http.StatusOK, "success")
}
