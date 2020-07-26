package helpers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func ResponseJSON(c int, data []string, context *gin.Context){
	context.JSON(
		http.StatusCreated,
		gin.H{
			"status":http.StatusCreated, 
			"code_error":c,
			"message":msgError[c], 
			"data":data})
}

func msgError(code int){
	var msg []string

	msg[45]="error username / password"
	msg[67]=""
	return msg[code]
}
