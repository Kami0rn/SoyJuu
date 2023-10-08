package entities

import (
	"gorm.io/gorm"

	"time"
)

type Salary struct {

	gorm.Model

	Amount float64 

	Point float32

	DateTime time.Time `gorm:"autoCreateTime:true" json:"dateTime"`

	//FK 
	HistoryID *uint
	History []History `gorm:"foreignKey:HistoryID"`

	RiderID *uint
	Rider []Rider `gorm:"foreignKey:RiderID"`

}