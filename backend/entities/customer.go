package entities


import (

"gorm.io/gorm"

)


type Customer struct {

	gorm.Model

	FirstName string

	LastName string

	Username string `gorm:"uniqueIndex"`

	Password string

	Address string

	Email string `gorm:"uniqueIndex"`

	Phone string

	Gender string

	//FK export
	Orders []Order `gorm:"foreignKey:CustomerID"`

	Deliveries []Delivery `gorm:"foreignKey:CustomerID"`

	Payments []Payment `gorm:"foreignKey:CustomerID"`

}



