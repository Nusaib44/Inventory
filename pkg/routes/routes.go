package routes

import (
	"inventory/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(g *gin.Engine) {

	g.GET("/listproduct", controllers.ListProducts)
	g.GET("/listproductbyid", controllers.ListProductByID)
	g.POST("/addproduct", controllers.AddProducts)
	g.PATCH("/editproduct", controllers.EditProduct)
	g.DELETE("/deleteproduct", controllers.DeleteProdduct)

	g.GET("/listprovider",controllers.ListProvidor)
	g.GET("/listproviderbyid",controllers.ListProvidorbyId)
	g.POST("/addprovider",controllers.AddProvidor)
	g.PATCH("/editprovider",controllers.EditProvidor)
	g.DELETE("/deleteprovider",controllers.DeleteProvidor)

	g.GET("/listorder")
	g.GET("/listorderbyid")
	g.POST("/addorder")
	g.PATCH("/cancelorder")
	g.PATCH("/editoredr")
	g.DELETE("/deleteorder")

	g.GET("/listreturn")
	g.GET("/listreturnbyid")
	g.POST("/addreturn")
	g.PATCH("/editreturn")

	g.GET("/listpayment_methods", controllers.ListPayment)
	g.GET("/listpayment_method_byid", controllers.ListPaymentById)
	g.POST("/addpayment_method", controllers.AddPayment)
	g.PATCH("/editpayment_method", controllers.EditPayment)
	g.DELETE("/deletepayment_method", controllers.DeletePaymentMethod)

}

////////////////stat////////////////
