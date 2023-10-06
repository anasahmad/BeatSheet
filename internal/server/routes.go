package server

import (
	"BeatSheet/internal/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func routing(router *gin.RouterGroup) {
	//health endpoints
	health(router)
	//beatSheet endpoints
	beatSheet(router)
	//beat endpoints
	beat(router)
	//act endpoints
	act(router)

	ai(router)
}

func health(router *gin.RouterGroup) {
	router.GET("/health", controller.Liveness)

	//prometheus metrics
	router.GET("/metrics")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func beatSheet(router *gin.RouterGroup) {
	router.POST("", controller.POSTBeatSheet)
	router.GET("/:id", controller.GETBeatSheet)
	router.PUT("/:id/", controller.PUTBeatSheet)
	router.DELETE("/:id", controller.DELETEBeatSheet)
	router.GET("", controller.GETBeatSheets)
}

func beat(router *gin.RouterGroup) {
	router.POST("/:id/beat", controller.POSTBeat)
	router.PUT("/:id/beat/:beatId", controller.PUTBeat)
	router.DELETE("/:id/beat/:beatId", controller.DELETEBeat)
}

func act(router *gin.RouterGroup) {
	router.POST("/:id/beat/:beatId/act", controller.POSTAct)
	router.PUT("/:id/beat/:beatId/act/:actId", controller.PUTAct)
	router.DELETE("/:id/beat/:beatId/act/:actId", controller.DELETEAct)
}

func ai(router *gin.RouterGroup) {
	router.GET("/predict/act", controller.GetSuggestedAct)
	router.GET("/predict/beat", controller.GetSuggestedBeat)
}
