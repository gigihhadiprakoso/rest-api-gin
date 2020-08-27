package main

import (
	// "fmt"
	"rest-api-gin/config"
	"github.com/gin-gonic/gin"
	// . "rest-api-gin/models"
	"rest-api-gin/controllers"
	// "rest-api-gin/seed"
	// "log"
)

var db = config.ConnectDB()

func main() {
	// seed.Load(db);
	router := gin.Default()

	product := router.Group("product")
	{
		product.POST("add",controllers.AddProduct)
		product.GET("",controllers.FindProduct)
		product.GET(":id",controllers.FindProductByID)
		product.DELETE(":id",controllers.DeleteProduct)
		product.PUT(":id",controllers.EditProduct)
	}

	/*category := router.Group("category")
	{
		category.POST("add",controllers.AddCategory)
		category.GET("",controllers.FindCategory)
	}*/


	router.POST("register",controllers.Register)
	router.POST("login",controllers.Login)

	router.Run()
}