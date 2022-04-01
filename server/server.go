package server

import (
	"GoWebServer/config"
	"net/http"
	"time"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Init() {
	config := config.Config()
	router := NewRouter()
	server := &http.Server{
		Addr:           config.GetString("server.addr"),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.ListenAndServe()
}
