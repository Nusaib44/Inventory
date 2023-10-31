package models

type ProductListing struct {
	Name             string
	Category         int
	Brand            string
	ShortDescription string
	Price            int64
	MainImg          string
	Provider         string
}
type merchentListing struct {
	Name        string
	CompanyName string
	Pincode     int
	Email       string
	PhoneNumber string
}
type ProductByID struct {
	ID               int
	Name             string
	Category         string
	Brand            string
	ShortDescription string
	Description      string
	Price            int64
	WholesalePrice   int64
	Stock            int64
	TotalOrder       int
	StockPurchased   int
	Revenue          int
	Provider         merchentListing
	WhareHouse       string
	MainImg          string
	Img2             string
	Img3             string
}
