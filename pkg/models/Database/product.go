package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name           string `json:"name"`
	Code           int    `json:"code"`
	Category       int    `gorm:"foreignkey"`
	Brand          string
	Description    string
	Price          int64 `json:"price"`
	WholesalePrice int64
	Stock          int64 `json:"stock"`
	TotalOrder     int
	StockPurchased int
	Profit         int64
	Provider       int `gorm:"foreignkey"`
	WhareHouse     int `gorm:"foreignkey"`
	MainImg        string
	Img2           string
	Img3           string
}

type Order struct {
	gorm.Model
	ProductId   int `gorm:"foreignkey"`
	UserId      int `gorm:"foreignkey"`
	Address     int `gorm:"foreignkey"`
	Quantity    int
	TotalPrice  int
	PaymentType int `gorm:"foreignkey"`
	Pincode     int
	Status      string
}

type Return struct {
	gorm.Model
	OrderId    int `gorm:"foreignkey"`
	Quantity   int
	TotalPrice int
	PaymentType   string
	Status     string
}


