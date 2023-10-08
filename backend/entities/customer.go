package entities


import (

"gorm.io/gorm"

)


type Customer struct {

	gorm.Model

	FirstName string `gorm:"not null"`

	LastName string `gorm:"not null"`

	Username string `gorm:"uniqueIndex"`

	Password string `gorm:"not null"`

	Address string `gorm:"not null"`

	Email string `gorm:"uniqueIndex"`

	Phone string `gorm:"not null"`

	Gender string `gorm:"not null"`

	//FK export
	Orders []Order `gorm:"foreignKey:CustomerID"`

	Deliveries []Delivery `gorm:"foreignKey:CustomerID"`

	Payments []Payment `gorm:"foreignKey:CustomerID"`

}



