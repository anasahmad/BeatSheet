package controller

import (
	"BeatSheet/internal/ai"
	"github.com/gin-gonic/gin"
)

func GetSuggestedBeat(ctx *gin.Context) {
	beat, err := ai.NewAIModel().GetNextBeatSuggestion()
	responseHandler(ctx, beat, err.Error())
}

func GetSuggestedAct(ctx *gin.Context) {
	act, err := ai.NewAIModel().GetNextActSuggestion()
	responseHandler(ctx, act, err.Error())
}
