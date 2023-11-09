package controllers

import (
	"encoding/json"
	db "inventory/pkg/db"
	models "inventory/pkg/models/Database"
	Response "inventory/pkg/models/Response"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// list product
func ListProducts(g *gin.Context) {

	var products []Response.ProductListing

	pagestring := g.Query("page")
	PageNumber, _ := strconv.Atoi(pagestring)
	offset := (PageNumber - 1) * 10
	db.DB.Raw("SELECT * FROM products ORDER BY name OFFSET ? ", offset).Scan(&products)
	Surcessmessage(g, "displaying products", products)
}

// list product by id
func ListProductByID(g *gin.Context) {

	var product models.Product

	params := g.Query("id")
	id, _ := strconv.Atoi(params)

	db.DB.Find(&product, id)
	if product.ID < 1 {
		ErrorMessage(g, 200, "please enter a valid product id", "product not found", product)
		return
	}
	Surcessmessage(g, "displaying product", product)

}

// add product
func AddProducts(g *gin.Context) {

	var product models.Product
	var homepic string

	// taking product information
	body := g.PostForm("product")
	err := json.Unmarshal([]byte(body), &product)
	if err != nil {
		ErrorMessage(g, 200, "failed to unmarshal the body", err.Error(), product)
	}

	// taking main img
	mainimg, err1 := g.FormFile("mainimg")
	if err1 != nil {
		ErrorMessage(g, 200, "failed to get main image", err1.Error(), mainimg)
		return
	} else {
		exrention := filepath.Ext(mainimg.Filename)
		image := "product" + uuid.New().String() + exrention
		homepic = image
		g.SaveUploadedFile(mainimg, "./public/"+image)
		product.MainImg = homepic
	}

	//taking imag2
	file := g.Request.MultipartForm.File["image2"]
	for _, img := range file {
		extention := filepath.Ext(img.Filename)
		image := "product" + uuid.New().String() + extention
		g.SaveUploadedFile(img, "./public/"+image)
		if product.Img2 == "" {
			product.Img2 = image
		} else {
			product.Img3 = image
		}

	}
	//adding to db
	result := db.DB.Create(&product)

	if result.Error != nil {
		ErrorMessage(g, 502, "Failed to add product", result.Error.Error(), product)
		return
	}
	Surcessmessage(g, "Product added Surcessfully", product)

}

// edit product
func EditProduct(g *gin.Context) {

	var product models.Product

	param := g.Query("id")
	id, _ := strconv.Atoi(param)
	println(id, "-----------------------///////////////////-------------")

	if err := g.Bind(&product); err != nil {
		ErrorMessage(g, 200, "error while binding", err.Error(), product)
	}
	db.DB.Model(&product).Where("ID=?", id).Updates(product)
	db.DB.Raw("SELECT * FROM products WHERE id=?", id).Scan(&product)

	Surcessmessage(g, "product edited surcessfully", product)
}

// ddelete product
func DeleteProdduct(g *gin.Context) {

	var product models.Product

	param := g.Query("id")
	id, _ := strconv.Atoi(param)

	db.DB.Raw("SELECT *FROM products WHERE id=?", id).Scan(&product)
	db.DB.Raw("DELETE FROM products WHERE id=?", id).Scan(&product)

	Surcessmessage(g, "product deleted surcesfully", product)
}
