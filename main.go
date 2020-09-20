package main

import (
	"rest-api-gin/config"
	"github.com/gin-gonic/gin"
	. "rest-api-gin/models"
	"rest-api-gin/controllers"
	// "rest-api-gin/seed"
	. "rest-api-gin/helpers"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var db = config.ConnectDB()

func main() {
	// seed.Load(db);
	router := gin.Default()
	
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