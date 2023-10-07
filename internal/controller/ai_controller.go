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

// GetSuggestedBeat godoc
//
//	@Summary		GET Beat
//	@Description	GET GetSuggestedBeat
//	@Tags			Predict
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	model.Beat
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/predict/beat  [get]
func (c *AIController) GetSuggestedBeat(ctx *gin.Context) {
	beat, err := c.aiModel.GetNextBeatSuggestion()
	responseHandler(ctx, beat, err)
}

// GetSuggestedAct godoc
//
//	@Summary		GET Act
//	@Description	GET GetSuggestedAct
//	@Tags			Predict
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	model.Act
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/predict/act  [get]
func (c *AIController) GetSuggestedAct(ctx *gin.Context) {
	act, err := c.aiModel.GetNextActSuggestion()
	responseHandler(ctx, act, err)
}
