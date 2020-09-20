package controllers

import (
	"github.com/gin-gonic/gin"

	. "rest-api-gin/models"
	// "rest-api-gin/config"
	. "rest-api-gin/helpers"
	// "strconv"
)

var c []Categories

func AddCategory(ctx *gin.Context){
	name := ctx.PostForm("name")

	category := Categories{
		Name: name,
		CompanyID: ID_Company,
	}
	
	result := db.Create(&category); 
	if result.Error != nil{
		ResponseJSON(ctx, 998, result.Error)
	}else{
		res := make(map[string]uint)
		res["id"]=category.ID
		ResponseJSON(ctx, 100, res)
	}
}

func FindCategories(ctx *gin.Context){
	result := db.Find(&c,"company_id=?",ID_Company)

	if result.Error != nil {
		ResponseJSON(ctx, 998, result.Error)
	} else if result.RowsAffected == 0 {
		ResponseJSON(ctx, 994, "")
	} else {
		ResponseJSON(ctx, 100, result.Value)
	}
}

func FindCategoryByID(ctx *gin.Context){
	id := ctx.Param("id")
	result := db.Find(&c,id)

	if result.Error != nil {
		ResponseJSON(ctx, 998, result.Error)
	} else if result.RowsAffected == 0 {
		ResponseJSON(ctx, 994, "")
	} else {
		ResponseJSON(ctx, 100, result.Value)
	}

}

func EditCategory(ctx *gin.Context){
	id := ctx.Param("id")
	name := ctx.PostForm("name")

	result := db.Model(&c).Where(id).Updates(map[string]interface{}{
		"name":name,
	})

	if result.Error != nil {
		ResponseJSON(ctx, 998, result.Error)
	}else{
		ResponseJSON(ctx, 100, id)
	}
}

func DeleteCategory(ctx *gin.Context){
	id := ctx.Param("id")

	db.Model(&c).Where(id).Update("is_deleted",1)
	result := db.Where(id).Delete(&Categories{})
	if result.Error != nil {
		ResponseJSON(ctx, 998, result.Error)
	} else {
		ResponseJSON(ctx, 100, result)
	}
}