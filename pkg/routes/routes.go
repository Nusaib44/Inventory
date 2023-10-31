package routes

import (
	"inventory/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(g *gin.Engine) {

	g.GET("/listproduct", controllers.ListProducts)
	g.GET("/listproductbyid", controllers.ListProductByID)
	g.POST("/addproduct", controllers.AddProducts)
	g.PATCH("/editproduct")
	g.DELETE("/deleteproduct")

	g.GET("/listprovider")
	g.GET("/listproviderbyid")
	g.POST("/addprovider")
	g.PATCH("/editprovider")
	g.DELETE("/deleteprovider")

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

	g.GET("/listpayment_methods")
	g.GET("/listpayment_method_byid")
	g.POST("/addpayment_method")
	g.PATCH("/editpayment_method")
	g.DELETE("/deletepayment_method")

}

////////////////stat////////////////
