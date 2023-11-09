package controllers

import (
	"inventory/pkg/db"
	models "inventory/pkg/models/Database"
	"strconv"

	"github.com/gin-gonic/gin"
)

var payment models.Payment
var new models.Payment

// ///////////
func AddPayment(g *gin.Context) {

	if err := g.Bind(&payment); err != nil {
		ErrorMessage(g, 200, "error while binding", err.Error(), payment)
	}

	if result := db.DB.Create(&payment); result.Error != nil {
		ErrorMessage(g, 500, "error while binding", result.Error.Error(), payment)
	}

	Surcessmessage(g, "payment added surcessfully", payment)
	payment = new

}

// ////////////
func EditPayment(g *gin.Context) {

	params := g.Query("id")
	id, _ := strconv.Atoi(params)

	if err := g.Bind(&payment); err != nil {
		ErrorMessage(g, 200, "error while binding", err.Error(), payment)
	}

	if result := db.DB.Model(&payment).Where("id=?", id).Updates(payment).Scan(&payment); result.Error != nil {
		ErrorMessage(g, 501, "Can't edit payment method", result.Error.Error(), payment)
	}
	db.DB.Raw("SELECT *FROM payments WHERE id=?", id).Scan(&payment)
	Surcessmessage(g, "payment method edited surcessfuly", payment)
	payment = new

}

func DeletePaymentMethod(g *gin.Context) {

	param := g.Query("id")
	id, _ := strconv.Atoi(param)

	db.DB.Raw("SELECT *FROM payments WHERE id=?", id).Scan(&payment)
	db.DB.Raw("DELETE FROM payments WHERE id=?", id).Scan(&payment)

	Surcessmessage(g, "paymentt deleted surcesfully", payment)
	payment = new
}

func ListPayment(g *gin.Context) {
	var payments []models.Payment
	pagestring := g.Query("page")
	PageNumber, _ := strconv.Atoi(pagestring)
	offset := (PageNumber - 1) * 10
	db.DB.Raw("SELECT * FROM payments ORDER BY type OFFSET ? ", offset).Scan(&payments)
	Surcessmessage(g, "displaying payment methods", payments)
	payment = new

}

func ListPaymentById(g *gin.Context) {
	params := g.Query("id")
	id, _ := strconv.Atoi(params)

	db.DB.Find(&payment, id)
	println(payment.ID, "--------------------")
	if payment.ID < 1 {
		ErrorMessage(g, 200, "Please enter a valid payment id", "Payment Method not found", payment)
		return
	}
	Surcessmessage(g, "Displaying payment by id", payment)
	payment = new
}
