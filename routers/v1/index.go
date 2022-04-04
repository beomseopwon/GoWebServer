package v1

import (
	"github.com/gin-gonic/gin"
)

func BindRoutes(r *gin.RouterGroup) {
	BindHealthRoutes(r)
	BindTokenRoutes(r)
}
