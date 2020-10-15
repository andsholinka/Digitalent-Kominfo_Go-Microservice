package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andsholinka/Digitalent-Kominfo_Go-Microservice/menu-service/config"
	"github.com/andsholinka/Digitalent-Kominfo_Go-Microservice/menu-service/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/andsholinka/Digitalent-Kominfo_Go-Microservice/menu-service/handler"
	"github.com/gorilla/mux"
)

func main() {
	cfg := config.Config{
		Database: config.Database{
			Driver:   "mysql",
			Host:     "localhost",
			Port:     "3306",
			User:     "root",
			Password: "admin",
			DbName:   "digitalent_microservice",
			Config:   "charset=utf8&parseTime=True&loc=Local",
		},
	}

	db, err := initDB(cfg.Database)
	if err != nil {
		log.Panic(err)
		return
	}

	router := mux.NewRouter()

	menuHandler := handler.MenuHandler{
		Db: db,
	}

	router.Handle("/add-menu", http.HandlerFunc(menuHandler.AddMenu))
	router.Handle("/menu", http.HandlerFunc(menuHandler.GetMenu))

	fmt.Println("Menu service listen on port :8000")
	log.Panic(http.ListenAndServe(":8000", router))
}

func initDB(dbConfig config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.Config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(database.Menu{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connection to DB")

	return db, nil
}
