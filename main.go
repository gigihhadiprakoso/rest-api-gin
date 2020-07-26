package main

import (
	"rest-api-gin/config"
	"github.com/gin-gonic/gin"
	"rest-api-gin/models"
	"rest-api-gin/controllers"
)

var db = config.ConnectDB()

func init(){
	db.AutoMigrate(&models.Products{})
}

func main() {
	router := gin.Default()

	v1 := router.Group("api/v1/")
	
	{
		v1.POST("add",controllers.CreateProduct)
		v1.GET("",controllers.GetProducts)
		// v1.GET(":id",getProductById)
		// v1.DELETE(":id",deleteProductById)
		// v1.PUT(":id",updateProductById)
	}

	router.Run()
}