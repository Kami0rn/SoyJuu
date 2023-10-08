package entities

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {

	gorm.Model

	Status string 

	TotalPrice float32

	CreatedAt time.Time	`gorm:"autoCreateTime:true"`

	//FK
	CustomerID *uint
	Customer Customer `gorm:"foreignKey:CustomerID"`

	PaymentID *uint
	Payment Payment `gorm:"foreignKey:PaymentID"`

	//FK export
	OrderFoods []OrderFood `gorm:"foreignKey:OrderID"`

	Deliveries []Delivery `gorm:"foreignKey:OrderID"`

	Payments []Payment `gorm:"foreignKey:OrderID"`

}


//https://github.com/edwindvinas/shopping-cart-api/blob/master/cart/ShoppingCart.go