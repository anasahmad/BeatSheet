package server

import "github.com/gin-gonic/gin"

type Definition struct {
	Address     string
	ContextRoot string
}

var (
	engine     *gin.Engine
	definition = Definition{
		Address:     ":8080",
		ContextRoot: "/v0/beatsheet",
	}
)

func Serve() {
	engine = gin.Default()
	router := engine.Group(definition.ContextRoot)
	routing(router)

	err := engine.Run(definition.Address)

	if err != nil {
		return
	}
}
