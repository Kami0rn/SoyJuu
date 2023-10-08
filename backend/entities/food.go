package entities

import (

	"gorm.io/gorm"
)

type Food struct {
		gorm.Model

		FoodName string 

		FoodPrice float32 

		Description string 

		Path string 


		//FK export
		OrderFoods []OrderFood `gorm:"foreignKey:FoodID"`

		FoodMenus []FoodMenu `gorm:"foreignKey:FoodID"`
	}
