package controllers

import (
	// "net/http"
	"github.com/gin-gonic/gin"

	. "rest-api-gin/models"
	// "rest-api-gin/config"
	. "rest-api-gin/helpers"
	"strconv"
)

var p []Products

func AddProduct(ctx *gin.Context){
	name := ctx.PostForm("name")
	category, _ := strconv.ParseUint(ctx.PostForm("id_category"),10,32)
	brand, _ := strconv.ParseUint(ctx.PostForm("id_brand"),10,32)

	product := Products{
		Name: name, 
		CategoryID: uint(category),
		BrandID: uint(brand),
		CompanyID: ID_Company,
	}

	result := db.Create(&product); 
	if result.Error != nil{
		ResponseJSON(ctx, 998, result.Error)
	}else{
		res := make(map[string]uint)
		res["id"]=product.ID
		ResponseJSON(ctx, 100, res)
	}
}

func FindProduct(ctx *gin.Context) {
	result := db.Find(&p,"company_id=?",ID_Company)

	if result.Error != nil {
		ResponseJSON(ctx, 998, result.Error)
	} else {
		ResponseJSON(ctx, 100, result.Value)
	}
}

func FindProductByID(ctx *gin.Context){
	id := ctx.Param("id")
	result := db.Find(&p,id)

	if result.Error != nil {
		ResponseJSON(ctx, 998, result.Error)
	} else {
		ResponseJSON(ctx, 100, result.Value)
	}

}

func EditProduct(ctx *gin.Context){
	id := ctx.Param("id")
	// name := c.PostForm("name")
	id_category,_ := strconv.ParseUint(ctx.PostForm("id_category"),10,32)
	id_brand,_ := strconv.ParseUint(ctx.PostForm("id_brand"),10,32)

	// db.First(&p, uint(id))
	result := db.Model(&p).Where(id).Updates(map[string]interface{}{
		"category_id":uint(id_category),
		"brand_id":uint(id_brand),
	})

	if result.Error != nil {
		ResponseJSON(ctx, 998, result.Error)
	}else{
		ResponseJSON(ctx, 100, id)
	}
}

func DeleteProduct(ctx *gin.Context){
	id := ctx.Param("id")

	db.Model(&p).Where(id).Update("is_deleted",1)
	result := db.Where(id).Delete(&Products{})
	if result.Error != nil {
		ResponseJSON(ctx, 998, result.Error)
	} else {
		ResponseJSON(ctx, 100, result)
	}
}