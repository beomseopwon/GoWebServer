package v1

import (
	"github.com/gin-gonic/gin"
)

func BindRoutes(r *gin.Engine) {
	rootGroup := r.Group("/v1")
	BindHealthRoutes(rootGroup)
	tokenGroup := rootGroup.Group("/token")
	BindTokenRoutes(tokenGroup)
}
