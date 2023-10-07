package app

import (
	"BeatSheet/internal/ai"
	"BeatSheet/internal/controller"
	"BeatSheet/internal/data"
	"BeatSheet/internal/server"
	"context"
)

type app struct {
	actController       controller.ActController
	beatController      controller.BeatController
	beatSheetController controller.BeatSheetController
	aiController        controller.AIController
	healthController    controller.HealthController
	server              server.Server
}

func NewApp() app {
	client := data.NewMongoClient(context.TODO())
	aiModel := ai.NewAIModel()
	app := app{
		beatSheetController: controller.NewBeatSheetController(client),
		beatController:      controller.NewBeatController(client, aiModel),
		actController:       controller.NewActController(client, aiModel),
		aiController:        controller.NewAIController(client, aiModel),
		healthController:    controller.NewHealthController(),
		server:              server.NewServer(":8080", "/v0/beatsheet"),
	}
	return app
}

func (a *app) Run() {
	a.server.Serve(a.beatSheetController, a.beatController, a.actController, a.aiController, a.healthController)
}
