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

// func init(){
	// db.AutoMigrate(&models.Products{})

	/*err := db.Debug().DropTableIfExists(&Categories{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&Categories{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range seed.Scategories {
		fmt.Println(&seed.Scategories[i])
		// err = db.Debug().Model(&Categories{}).Create(seed.Scategories[i]).Error
		if err != nil {
			log.Fatalf("cannot seed categories table: %v", err)
		}
	}*/

	// 	err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
	// 	if err != nil {
	// 		log.Fatalf("cannot seed posts table: %v", err)
	// 	}
	// }
// }

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

	router.Run()
}