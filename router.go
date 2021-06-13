package main

import (
	"net/http"

	_"gin-api-boilerplate/src/apps/examples"
	"gin-api-boilerplate/src/middlewares"
	"gin-api-boilerplate/src/utils"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func APIRouter() *gin.Engine {
	r := gin.Default()

	//*Use the middlewares CORS & Gzip
	r.Use(middlewares.CORSMiddleware())
	r.Use(gzip.Gzip(gzip.BestCompression))

	//*Handling docs with swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//*Handle if no routes existed
	r.NoRoute(func(c *gin.Context) {
		utils.ResponseFormatter(http.StatusNotFound, "Not Found.", nil, nil, c)
	})

	//*Create new controllers here


	//*Create supergroup for versioning api
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", middlewares.RequestIDMiddleware(), func(c *gin.Context) {
			utils.ResponseFormatter(http.StatusOK, "Pong", nil, nil, c)
		})

	}

	return r
}
