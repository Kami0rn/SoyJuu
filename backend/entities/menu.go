package entities

import (
	"gorm.io/gorm"
)

type Menu struct {

	gorm.Model

	Title string 

	Description string 
	//FK export
	FoodMenus []FoodMenu `gorm:"foreignKey:MenuID"`
}