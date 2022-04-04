package v1

import (
	v1 "GoWebServer/controllers/v1"

	"github.com/gin-gonic/gin"
)

func BindTokenRoutes(router *gin.RouterGroup) {
	router.POST("/balanceof", v1.BalanceOf)
	router.POST("/tokesnofowner", v1.TokensOfOwner)
	router.POST("/tokenuri", v1.TokenURI)
}
