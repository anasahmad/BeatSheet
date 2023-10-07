package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthController struct {
}

func NewHealthController() HealthController {
	return HealthController{}
}
func (c *HealthController) Liveness(context *gin.Context) {
	context.Status(http.StatusOK)
}
