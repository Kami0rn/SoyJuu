package entities

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("soy-juu.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(
		&Customer{},
		&Delivery{}, 
		&Food{}, 
		&FoodMenu{}, 
		&History{}, 
		&Menu{}, 
		&Order{}, 
		&OrderFood{}, 
		&Rider{},
		&Payment{},
		&Salary{},
		&Wallet{})

	db = database
}