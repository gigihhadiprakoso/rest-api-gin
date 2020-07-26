package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"rest-api-gin/models"
	"rest-api-gin/config"
	"rest-api-gin/helpers"
)

var db = config.ConnectDB()

func CreateProduct(c *gin.Context){
	product := models.Products{Name: c.PostForm("nama")}
	db.Create(&product)
	c.JSON(http.StatusCreated,gin.H{"status":http.StatusCreated, "message":"Product Created Successfully!", "id":product.ID})
}

func GetProducts(c *gin.Context) {
	
}