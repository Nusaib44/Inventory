package routes

import (
	"inventory/controllers/employe"
	"github.com/gin-gonic/gin"
)

func Routes(g *gin.Engine) {

	g.GET("/listproduct",controllers.ListProducts)
	g.GET("/listproductbyid")
	g.GET("/listtrending")
	g.GET("/list similar")
	// g.GET("/")
	// g.GET("/")
	// g.GET("/")
	// g.GET("/")
	// g.GET("/")
	// g.GET("/")

}

////////////////stat////////////////
