package controllers

import (
	"inventory/pkg/db"
	models "inventory/pkg/models/Database"
	Response "inventory/pkg/models/Response"
	"strconv"

	"github.com/gin-gonic/gin"
)

var providor models.Providor
var newProvidor models.Providor

func ListProvidor(g *gin.Context) {
	var providors []Response.ProvidorListing

	// pagestring := g.Query("page")
	// PageNumber, _ := strconv.Atoi(pagestring)
	// offset := (PageNumber - 1) * 10
	db.DB.Raw("SELECT * FROM providors ORDER BY name ").Scan(&providors)
	Surcessmessage(g, "displaying providors", providors)

}

func ListProvidorbyId(g *gin.Context) {
	params := g.Query("id")
	id, _ := strconv.Atoi(params)

	db.DB.Find(&providor, id)
	if providor.ID < 1 {
		ErrorMessage(g, 200, "please enter a valid providor id", "providor not found", providor)
		return
	}
	Surcessmessage(g, "displaying providor", providor)
	providor = newProvidor

}

func AddProvidor(g *gin.Context) {

	if err := g.Bind(&providor); err != nil {
		ErrorMessage(g, 200, "error while binding", err.Error(), providor)
	}

	if result := db.DB.Create(&providor); result.Error != nil {
		ErrorMessage(g, 500, "error while binding", result.Error.Error(), providor)
	}

	Surcessmessage(g, "providor added surcessfully", providor)
	providor = newProvidor
}

func EditProvidor(g *gin.Context) {
	params := g.Query("id")
	id, _ := strconv.Atoi(params)

	if err := g.Bind(&providor); err != nil {
		ErrorMessage(g, 200, "error while binding", err.Error(), providor)
	}

	if result := db.DB.Model(&providor).Where("id=?", id).Updates(providor).Scan(&providor); result.Error != nil {
		ErrorMessage(g, 501, "Can't edit providor method", result.Error.Error(), providor)
	}
	db.DB.Raw("SELECT *FROM providors WHERE id=?", id).Scan(&providor)
	Surcessmessage(g, "providor details edited surcessfuly", providor)
	providor = newProvidor
}

func DeleteProvidor(g *gin.Context) {
	param := g.Query("id")
	id, _ := strconv.Atoi(param)

	db.DB.Raw("SELECT *FROM providors WHERE id=?", id).Scan(&providor)
	db.DB.Raw("DELETE FROM providors WHERE id=?", id).Scan(&providor)

	Surcessmessage(g, "providort deleted surcesfully", providor)
	providor = newProvidor
}
