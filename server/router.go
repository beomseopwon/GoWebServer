package server

import (
	v1 "GoWebServer/routers/v1"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string { // 커스텀 로그 (아파치에서 출력하는 형식)
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.UTC.String()),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routerGroupV1 := router.Group("/v1")
	v1.InitRoutes(routerGroupV1)

	return router
}
