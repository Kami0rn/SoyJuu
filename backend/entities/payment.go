package entities

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {		

	gorm.Model

	IsPaid bool `gorm:"not null"`

	Method string `gorm:"not null"`

	CreatedAt time.Time `gorm:"autoCreateTime:true"`

	//FK

	OrderID *uint 

	CustomerID *uint 


	//FK export
	Orders []Order `gorm:"foreignKey:PaymentID"`

	Wallets []Wallet `gorm:"foreignKey:PaymentID"`

	PayAtDeliveries []PayAtDelivery `gorm:"foreignKey:PaymentID"`

	Banks []Bank `gorm:"foreignKey:PaymentID"`


}

type Wallet struct {

	gorm.Model

	Remain float64 `gorm:"not null"`

	//FK

	PaymentID *uint
	Payment []Payment `gorm:"foreignKey:PaymentID"`

	//FK export
	TopUps []TopUp `gorm:"foreignKey:WalletID"`

}

type TopUp struct {

	gorm.Model

	TopUpAmount float64 `gorm:"not null"`

	//FK
	WalletID *uint
	Wallet []Wallet `gorm:"foreignKey:WalletID"`

}

type PayAtDelivery struct {

	gorm.Model

	//FK
	PaymentID *uint
	Payment []Payment `gorm:"foreignKey:PaymentID"`

}

type Card struct {

	gorm.Model

	CardNumber string `gorm:"not null"`

	CardHolder string `gorm:"not null"`

	ExpiredDate string `gorm:"not null"`

	AuditNumber string `gorm:"not null"`

	CardType string `gorm:"not null"`

}

type Bank struct {

	gorm.Model

	BankName string `gorm:"not null"`

	AccountNumber string `gorm:"not null"`

	AccountHolder string `gorm:"not null"`

	//FK
	PaymentID *uint
	Payment []Payment `gorm:"foreignKey:PaymentID"`

}

