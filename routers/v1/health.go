package v1

import (
	v1 "GoWebServer/controllers/v1"

	"github.com/gin-gonic/gin"
)

func SetHealthRoutes(router *gin.RouterGroup) {
	router.GET("/health", v1.Status)
}
