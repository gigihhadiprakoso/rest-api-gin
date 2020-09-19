package helpers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "fmt"
)

func ResponseJSON(context *gin.Context, c int, data interface{}){
	context.AbortWithStatusJSON(http.StatusOK, gin.H{"status":http.StatusOK, "code_error":c, "message":msgError(c), "data":data})
}

func msgError(code int) string {
	var msg [999]string

	msg[100]="Success"
	msg[101]="haloo"
	msg[2]="error nama "
	msg[998]="Error"
	msg[997]="Username / password tidak tepat"
	msg[996]="Token API dibutuhkan"
	msg[995]="Token sudah kedaluwarsa. Silakan login lagi"
	return msg[code]
}
