package controllers

import (
	"github.com/gin-gonic/gin"
	. "rest-api-gin/models"
	. "rest-api-gin/helpers"
	// "fmt"
	// "os"
	// "reflect"
	"encoding/json"
)

var u []Users
var abc map[string]interface{}

func Register(ctx *gin.Context){
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	firstname := ctx.PostForm("first_name")
	lastname := ctx.PostForm("last_name")
	companyname := ctx.PostForm("company_name")

	passwordHashed,_ := HashPassword(password)

	user := Users{
		FirstName: firstname, 
		LastName: lastname,
		Username: username,
		Password: passwordHashed,
		CompanyName: companyname,
	}

	result := db.Create(&user); 
	if result.Error != nil{
		ResponseJSON(ctx, 998, result.Error)
	}else{
		res := make(map[string]string)
		res["username"]=username
		res["name"]=firstname+" "+lastname
		ResponseJSON(ctx, 100, res)
	}
}

func Login(ctx *gin.Context){
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	result := db.Where("username = ?", username).First(&u)

	if result.Error != nil{
		ResponseJSON(ctx, 998, result.Error)
	}

	resultString, _ := json.Marshal(&result.Value)
	var user []Users
	json.Unmarshal([]byte(resultString), &user)

	r := CheckPasswordHash(password, user[0].Password)

	if r == false{
		ResponseJSON(ctx, 997, "")
	}

	res := make(map[string]string)
	res["token"], _ = CreateToken(user[0].ID)
	ResponseJSON(ctx, 100, res)
}

