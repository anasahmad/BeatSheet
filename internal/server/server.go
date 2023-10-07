package server

import (
	"BeatSheet/internal/controller"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Address     string
	ContextRoot string
}

func NewServer(add string, contextRoot string) Server {
	return Server{
		Address:     add,
		ContextRoot: contextRoot,
	}
}

var (
	definition = Server{
		Address:     ":8080",
		ContextRoot: "/v0/beatsheet",
	}
)

func (s *Server) Serve(sheetController controller.BeatSheetController, beatController controller.BeatController, actController controller.ActController, aiController controller.AIController, healthController controller.HealthController) {
	engine := gin.Default()
	router := engine.Group(s.ContextRoot)
	routing(router, sheetController, beatController, actController, aiController, healthController)

	err := engine.Run(s.Address)

	if err != nil {
		return
	}
}
