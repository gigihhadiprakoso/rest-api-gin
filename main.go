package main

import (
	"rest-api-gin/config"
	"github.com/gin-gonic/gin"
	// . "rest-api-gin/models"
	"rest-api-gin/controllers"
	"rest-api-gin/seed"
	. "rest-api-gin/helpers"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
	// "os"
	"regexp"
)

var db = config.ConnectDB()

func main() {
	seed.Load(db);
	router := gin.Default()
	
	router.Use(inputValidationMidleware())
	router.POST("register",controllers.Register)
	router.POST("login",controllers.Login)

	router.Use(authMiddleware())

	product := router.Group("product")
	{
		product.POST("add",controllers.AddProduct)
		product.GET("",controllers.FindProducts)
		product.GET(":id",controllers.FindProductByID)
		product.DELETE(":id",controllers.DeleteProduct)
		product.PUT(":id",controllers.EditProduct)
	}

	category := router.Group("category")
	{
		category.POST("add",controllers.AddCategory)
		category.GET("",controllers.FindCategories)
		category.GET(":id",controllers.FindCategoryByID)
		category.PUT(":id",controllers.EditCategory)
		category.DELETE(":id",controllers.DeleteCategory)
	}

	brand := router.Group("brand")
	{
		brand.POST("add",controllers.AddBrand)
		brand.GET("",controllers.FindBrands)
		brand.GET(":id",controllers.FindBrandByID)
		brand.PUT(":id",controllers.EditBrand)
		brand.DELETE(":id",controllers.DeleteBrand)
	}

	// supplier := router.Group("supplier")
	// {
	// 	supplier.POST("add",controllers.AddSupplier)
	// 	supplier.GET("",controllers.FindSuppliers)
	// 	supplier.GET(":id",controllers.FindSupplierByID)
	// 	supplier.PUT(":id",controllers.EditSupplier)
	// 	supplier.DELETE(":id",controllers.DeleteSupplier)
	// }

	// customer := router.Group("customer")
	// {
	// 	customer.POST("add",controllers.AddCustomer)
	// 	customer.GET("",controllers.FindCustomers)
	// 	customer.GET(":id",controllers.FindCustomerByID)
	// 	customer.PUT(":id",controllers.EditCustomer)
	// 	customer.DELETE(":id",controllers.DeleteCustomer)
	// }

	purchase := router.Group("purchase")
	{
		purchase.POST("add",controllers.AddPurchase)
		purchase.GET("",controllers.FindPurchases)
		purchase.GET(":id",controllers.FindPurchaseByID)
		purchase.PUT(":id",controllers.EditPurchase)
		purchase.DELETE(":id",controllers.DeletePurchase)
	}

	router.Run()
}

func authMiddleware() gin.HandlerFunc{
	
	return func(ctx *gin.Context){
		headerAuthorization := ctx.Request.Header.Get("Authorization")
		token := ExtractToken(headerAuthorization)

		if token == "" {
			ResponseJSON(ctx, 996, "")
			return
		}

		claims := Claims{}
		tokenDecoded,_ := jwt.ParseWithClaims(token, &claims, func(tokenDecoded *jwt.Token) (interface{}, error) {
			return []byte("gigih-rupawan"), nil
		})

		if !tokenDecoded.Valid || claims.ExpiresAt < time.Now().Unix() {
			ResponseJSON(ctx, 995, "")
			return
		}

		controllers.ID_Company = claims.CompanyID
		ctx.Next()
	}
}

func inputValidationMidleware() gin.HandlerFunc{
	return func(ctx *gin.Context){
		re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
		fmt.Println("tanggal",re.MatchString(ctx.PostForm("date")))

		ctx.Next()
	}
}