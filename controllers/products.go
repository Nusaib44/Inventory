package controllers

import (
	db "inventory/pkg/db"
	models "inventory/pkg/models/Database"

	"github.com/gin-gonic/gin"
)

func ListProducts(g *gin.Context) {

	var products []models.Product
	db.DB.Find(&products)
	Surcessmessage(g, "displaying products", products)
}
