package controllers

import (
	"github.com/gin-gonic/gin"

	. "rest-api-gin/models"
	// "rest-api-gin/config"
	. "rest-api-gin/helpers"
	// "strconv"
)

var pc []Purchases

func FindPurchases(ctx *gin.Context){
	result := db.Find(&pc,"company_id=?",ID_Company)

	if result.Error != nil {
		ResponseJSON(ctx, 998, result.Error)
	} else if result.RowsAffected == 0 {
		ResponseJSON(ctx, 994, "")
	} else {
		ResponseJSON(ctx, 100, result.Value)
	}
}

func FindPurchaseByID(ctx *gin.Context){
	id := ctx.Param("id")
	result := db.Find(&pc,id)

	if result.Error != nil {
		ResponseJSON(ctx, 998, result.Error)
	} else if result.RowsAffected == 0 {
		ResponseJSON(ctx, 994, "")
	} else {
		ResponseJSON(ctx, 100, result.Value)
	}
}

func AddPurchase(ctx *gin.Context){
	// name := ctx.PostForm("name")
	// warehouse_id, _ := strconv.ParseUint(ctx.PostForm("id_warehouse"),10,32)
	// supplier_id, _ := strconv.ParseUint(ctx.PostForm("id_supplier"),10,32)
	// date := ctx.PostForm("date")

	// purchase := Purchases{
	// 	ReferenceNo: GenerateReference(),
	// 	Date: date,
	// 	WarehouseID: uint(warehouse_id),
	// 	SupplierID: uint(supplier_id),
	// 	CompanyID: ID_Company,
	// }

	// result := db.Create(&purchase); 
	// if result.Error != nil{
	// 	ResponseJSON(ctx, 998, result.Error)
	// }else{
	// 	res := make(map[string]uint)
	// 	res["id"]=purchase.ID
	// 	ResponseJSON(ctx, 100, res)
	// }
}

func EditPurchase(ctx *gin.Context){

}

func DeletePurchase(ctx *gin.Context){

}