package server

import (
	"BeatSheet/internal/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func routing(router *gin.RouterGroup, sheetController controller.BeatSheetController, beatController controller.BeatController, actController controller.ActController, aiController controller.AIController, healthController controller.HealthController) {
	//health endpoints
	health(router, healthController)
	//beatSheet endpoints
	beatSheet(router, sheetController)
	//beat endpoints
	beat(router, beatController)
	//act endpoints
	act(router, actController)

	ai(router, aiController)
}

func health(router *gin.RouterGroup, controller controller.HealthController) {
	router.GET("/health", controller.Liveness)

	//prometheus metrics
	router.GET("/metrics")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func beatSheet(router *gin.RouterGroup, controller controller.BeatSheetController) {
	router.POST("", controller.POSTBeatSheet)
	router.GET("/:id", controller.GETBeatSheet)
	router.PUT("/:id/", controller.PUTBeatSheet)
	router.DELETE("/:id", controller.DELETEBeatSheet)
	router.GET("", controller.GETBeatSheets)
}

func beat(router *gin.RouterGroup, controller controller.BeatController) {
	router.POST("/:id/beat", controller.POSTBeat)
	router.PUT("/:id/beat/:beatId", controller.PUTBeat)
	router.DELETE("/:id/beat/:beatId", controller.DELETEBeat)
}

func act(router *gin.RouterGroup, controller controller.ActController) {
	router.POST("/:id/beat/:beatId/act", controller.POSTAct)
	router.PUT("/:id/beat/:beatId/act/:actId", controller.PUTAct)
	router.DELETE("/:id/beat/:beatId/act/:actId", controller.DELETEAct)
}

func ai(router *gin.RouterGroup, controller controller.AIController) {
	router.GET("/predict/act", controller.GetSuggestedAct)
	router.GET("/predict/beat", controller.GetSuggestedBeat)
}
