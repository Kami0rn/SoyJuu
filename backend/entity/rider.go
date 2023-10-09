package entity

import (
	"gorm.io/gorm"
)

type Rider struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Phone     string


}