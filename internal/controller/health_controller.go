package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Liveness(context *gin.Context) {
	context.Status(http.StatusOK)
}
