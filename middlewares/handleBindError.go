package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(context *gin.Context) {
	context.Next()

	// for _, err := range context.Errors {
	// 	// log, handle, etc.
	// }

	context.JSON(http.StatusInternalServerError, "")
}
