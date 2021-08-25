package helpers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "fmt"
	"strconv"
	"time"
)

func ResponseJSON(context *gin.Context, c int, data interface{}){
	context.AbortWithStatusJSON(http.StatusOK, gin.H{"status":http.StatusOK, "code_error":c, "message":msgError(c), "data":data})
}

func msgError(code int) string {
	var msg [999]string

	msg[100]="Success"
	msg[994]="Data tidak ditemukan"
	msg[995]="Token sudah kedaluwarsa. Silakan login lagi"
	msg[996]="Token API dibutuhkan"
	msg[997]="Username / password tidak tepat"
	msg[998]="Error"
	return msg[code]
}

func GenerateReference() string{
	return "REF/"+strconv.FormatInt(time.Now().Unix(),10)
}
