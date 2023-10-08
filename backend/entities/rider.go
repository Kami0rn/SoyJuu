package entities

import (
	"gorm.io/gorm"
)

type Rider struct {

	gorm.Model

	RideFirstName string

	RideLastName string

	RideUsername string `gorm:"unique"`

	RidePassword string

	RideTel int

	RideEmail string `gorm:"unique"`

	CarNum string


	//FK export
	// Deliveries []Delivery `gorm:"foreignKey:RiderID"`

	// Salarys []Salary `gorm:"foreignKey:RiderID"`

}