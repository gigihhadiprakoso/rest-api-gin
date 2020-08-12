package controllers

import (
	// "net/http"
	"github.com/gin-gonic/gin"

	. "rest-api-gin/models"
	"rest-api-gin/config"
	. "rest-api-gin/helpers"
	"strconv"
	// "fmt"
	// "reflect"
)

var db = config.ConnectDB()
var p []Products

func AddProduct(c *gin.Context){
	name := c.PostForm("name")
	category, _ := strconv.ParseUint(c.PostForm("id_category"),10,32)
	brand, _ := strconv.ParseUint(c.PostForm("id_brand"),10,32)

	product := Products{
		Name: name, 
		CategoryID: uint(category),
		BrandID: uint(brand),
	}

	result := db.Create(&product); 
	if result.Error != nil{
		ResponseJSON(c, 998, result.Error)
	}else{
		res := make(map[string]uint)
		res["id"]=product.ID
		ResponseJSON(c, 100, res)
	}
}

func FindProduct(c *gin.Context) {
	result := db.Find(&p)

	if result.Error != nil {
		ResponseJSON(c, 998, result.Error)
	} else {
		ResponseJSON(c, 100, result.Value)
	}
}

func FindProductByID(c *gin.Context){
	id := c.Param("id")
	result := db.Find(&p,id)

	if result.Error != nil {
		ResponseJSON(c, 998, result.Error)
	} else {
		ResponseJSON(c, 100, result.Value)
	}

}

func EditProduct(c *gin.Context){
	id := c.Param("id")
	// name := c.PostForm("name")
	id_category,_ := strconv.ParseUint(c.PostForm("id_category"),10,32)
	id_brand,_ := strconv.ParseUint(c.PostForm("id_brand"),10,32)

	// db.First(&p, uint(id))
	result := db.Model(&p).Where(id).Updates(map[string]interface{}{
		"category_id":uint(id_category),
		"brand_id":uint(id_brand),
	})

	if result.Error != nil {
		ResponseJSON(c, 998, result.Error)
	}else{
		ResponseJSON(c, 100, id)
	}
}

func DeleteProduct(c *gin.Context){
	id := c.Param("id")

	db.Model(&p).Where(id).Update("is_deleted",1)
	result := db.Where(id).Delete(&Products{})
	if result.Error != nil {
		ResponseJSON(c, 998, result.Error)
	} else {
		ResponseJSON(c, 100, result)
	}
}