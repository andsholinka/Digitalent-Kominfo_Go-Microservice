package database

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Auth struct {
	ID       int    `gorm:"primary_key" json:"-"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
}

func ValidateAuth(token string, db *gorm.DB) (*Auth, error) {
	var auth Auth

	if err := db.Where(&Auth{Token: token}).First(&auth).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Errorf("Invalid Toke")
		}
	}
	return &auth, nil
}

func (auth *Auth) SignUp(db *gorm.DB) error {
	//sebuah fungsi untuk mengecek apakah username yg diinput sudah ada di db?
	//kalau tidak ada, maka return error, maka action db.Create dijalankan
	if err := db.Where(&Auth{Username: auth.Username}).First(auth).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(auth).Error; err != nil {
				return err
			}
		}
	} else {
		return errors.Errorf("Duplicate Username")
	}
	return nil
}

func (auth *Auth) Login(db *gorm.DB) (*Auth, error) {
	if err := db.Where(&Auth{Username: auth.Username, Password: auth.Password}).First(auth).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Errorf("incorect username/password")
		}
	}
	return auth, nil
}
