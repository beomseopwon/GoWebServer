package v1

import (
	v1 "GoWebServer/controllers/v1"

	"github.com/gin-gonic/gin"
)

func SetTokenRoutes(router *gin.RouterGroup) {
	router.POST("/balanceof", v1.BalanceOf)
}
