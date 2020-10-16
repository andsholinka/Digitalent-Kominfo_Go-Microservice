package database

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Auth struct {
	ID       int    `gorm:"primary_key" json:="-"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func (auth *Auth) SignUp(db *gorm.DB) error {
	//select * form auth where username="fadhl"
	if err := db.Where(&Auth{Username: auth.Username}).First(auth).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(auth).Error; err != nil {
				return err
			}
		}
		return err
	}
	return nil
}

func (auth *Auth) Login(db *gorm.DB) (*Auth, error) {
	if err := db.Where(&Auth{Username: auth.Username, Password: auth.Password}).First(auth).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Errorf("incorect email/password")
		}
	}
	return auth, nil
}
