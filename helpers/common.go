package helpers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "fmt"
)

func ResponseJSON(ctx *gin.Context, c int, data interface{}){
	ctx.JSON(http.StatusOK, gin.H{"status":http.StatusOK, "code_error":c, "message":msgError(c), "data":data})
}

func msgError(code int) string {
	var msg [999]string

	msg[100]="Success"
	msg[101]="haloo"
	msg[2]="error nama "
	msg[998]="Error"
	return msg[code]
}
