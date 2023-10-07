package controller

import (
	"BeatSheet/internal/ai"
	"BeatSheet/internal/data"
	"github.com/gin-gonic/gin"
)

type AIController struct {
	dataClient data.DataClient
	aiModel    *ai.AIModel
}

func NewAIController(dataClient data.DataClient, aiModel *ai.AIModel) AIController {
	return AIController{dataClient: dataClient, aiModel: aiModel}
}

func (c *AIController) GetSuggestedBeat(ctx *gin.Context) {
	beat, err := c.aiModel.GetNextBeatSuggestion()
	responseHandler(ctx, beat, err)
}

func (c *AIController) GetSuggestedAct(ctx *gin.Context) {
	act, err := c.aiModel.GetNextActSuggestion()
	responseHandler(ctx, act, err)
}
