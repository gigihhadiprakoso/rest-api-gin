package controllers

import ( 
	"github.com/gin-gonic/gin"
	. "rest-api-gin/models"
	. "rest-api-gin/helpers"
)

var b []Brands

func FindBrands(ctx *gin.Context){
	result := db.Find(&b,"company_id=?",ID_Company)

	if result.Error != nil {
		ResponseJSON(ctx, 998, result.Error)
	} else if result.RowsAffected == 0 {
		ResponseJSON(ctx, 994, "")
	} else {
		ResponseJSON(ctx, 100, result.Value)
	}
}

func FindBrandByID(ctx *gin.Context){
	id := ctx.Param("id")
	result := db.Find(&b,id)

	if result.Error != nil {
		ResponseJSON(ctx, 998, result.Error)
	} else if result.RowsAffected == 0 {
		ResponseJSON(ctx, 994, "")
	} else {
		ResponseJSON(ctx, 100, result.Value)
	}
}

func AddBrand(ctx *gin.Context){
	name := ctx.PostForm("name")

	brand := Brands{
		Name: name,
		CompanyID: ID_Company,
	}
	
	result := db.Create(&brand); 
	if result.Error != nil{
		ResponseJSON(ctx, 998, result.Error)
	}else{
		res := make(map[string]uint)
		res["id"]=brand.ID
		ResponseJSON(ctx, 100, res)
	}
}