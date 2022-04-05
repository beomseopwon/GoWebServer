package v1

import (
	v1 "GoWebServer/controllers/v1"

	"github.com/gin-gonic/gin"
)

func BindContractRoutes(router *gin.RouterGroup) {
	router.POST("/contractaddress", v1.ContractAddress)
	router.POST("/mintnft", v1.MintNFT)
}
