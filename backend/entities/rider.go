package entities

import (
	"gorm.io/gorm"
)

type Rider struct {

	gorm.Model

	RideFirstName string `gorm:"not null"`

	RideLastName string `gorm:"not null"`

	RideUsername string `gorm:"unique;not null"`

	RidePassword string `gorm:"not null"`

	RideTel int `gorm:"not null"`

	RideEmail string `gorm:"uniqueIndex;check:email LIKE '%.com'"`

	DriverLicense int `gorm:"not null"`

	CarNum string `gorm:"not null"`


	//FK export
	Deliveries []Delivery `gorm:"foreignKey:RiderID"`

	Salarys []Salary `gorm:"foreignKey:RiderID"`

}