package controllers

import (
	// "net/http"
	"github.com/gin-gonic/gin"

	. "rest-api-gin/models"
	// "rest-api-gin/config"
	. "rest-api-gin/helpers"
	"strconv"
	// "fmt"
	// "os"
	"strings"

	"github.com/jinzhu/gorm"
)

var p []Products

func AddProduct(ctx *gin.Context){
	var wh_id uint
	var stock int
	warehouse_product :=make([]map[string]interface{},100)

	name := ctx.PostForm("name")
	category, _ := strconv.ParseUint(ctx.PostForm("id_category"),10,32)
	brand, _ := strconv.ParseUint(ctx.PostForm("id_brand"),10,32)
	unit,_ := strconv.ParseUint(ctx.PostForm("id_unit"),10,32)

	product := Products{
		Name: name, 
		CategoryID: uint(category),
		BrandID: uint(brand),
		UnitID:uint(unit),
		CompanyID: ID_Company,
	}

	db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Create(&product).Error; err != nil {
			return err
		}
		
		i := 0
		for key, value := range ctx.Request.PostForm {
			if strings.Index(key, "qty") != -1 {
				wh := strings.Split(key,"__")
				abc,_ := strconv.ParseUint(wh[1],10,32)
				wh_id = uint(abc)
				stock,_ = strconv.Atoi(value[0])

				warehouse_product[i] = map[string]interface{}{
					"Name":name,
					"WarehouseID":wh_id,
					"ProductID":product.ID,
					"Stock":stock,
				}
				i++
			}
		}

		if err := tx.Table("warehouse_products").Create(&warehouse_product).Error; err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	// if result.Error != nil{
	// 	ResponseJSON(ctx, 998, result.Error)
	// }else{
	// 	res := make(map[string]uint)
	// 	res["id"]=product.ID
	// 	ResponseJSON(ctx, 100, res)
	// }
}

func FindProducts(ctx *gin.Context) {
	result := db.Find(&p,"company_id=?",ID_Company)

	if result.Error != nil {
		ResponseJSON(ctx, 998, result.Error)
	} else if result.RowsAffected == 0 {
		ResponseJSON(ctx, 994, "")
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