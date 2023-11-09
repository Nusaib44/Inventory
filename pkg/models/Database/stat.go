package models

import "gorm.io/gorm"

type Providor struct {
	gorm.Model
	Name        string
	CompanyName string
	Pincode     int    `gorm:"foreignkey"`
	Email       string `gorm:"unique"`
	PhoneNumber string `gorm:"unique"`
	Paid        int
	Due         int
	Order       int
	Return      int
	Cancel      int
}

type Area struct {
	Pincode int `gorm:"unique"`
	Name    string
	State   string
	Distric string
	Country string
	Orders  string
	Returns string
	Cancels string
}

type WareHouse struct {
}
type Payment struct {
	gorm.Model
	Type   string
	Counnt int
}
type User struct {
	gorm.Model
	Username    string
	Email       string `gorm:"unique"`
	PhoneNumber string
}
type Address struct {
	gorm.Model
	UserId      int
	HouseName   string
	Street      string
	AddressLine string
	Pincode     int `gorm:"foreignkey"`
}
