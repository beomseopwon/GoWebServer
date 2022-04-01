package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Status godoc
// @Summary      서버 상태 체크
// @Description  서버가 살아 있는 확인
// @Tags         Status
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /v1/health [get]
func Status(context *gin.Context) {
	context.String(http.StatusOK, "Working!")
}
