package controller

import (
	"BeatSheet/internal/ai"
	"BeatSheet/internal/data"
	"BeatSheet/internal/model"
	"github.com/gin-gonic/gin"
)

// POSTBeat godoc
//
//	@Summary		Post Beat
//	@Description	Post Beat from the object in the body
//	@Tags			Beat
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"BeatSheet ID"
//	@Param			beat	body		model.Beat	true	"Beat Data"
//	@Success		200		{string}	id
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/beatsheet/:id/beat  [post]
func POSTBeat(ctx *gin.Context) {
	var beat model.Beat
	id := ctx.Param("id")
	mClient := data.NewMongoClient(ctx)
	ctx.ShouldBindJSON(&beat)
	ai.NewAIModel().AddBeat(beat)
	beatId, err := mClient.AddBeat(id, beat)
	responseHandler(ctx, beatId, err)
}

// PUTBeat godoc
//
//	@Summary		PUT Beat
//	@Description	PUT Beat from the object in the body
//	@Tags			Beat
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"BeatSheet ID"
//	@Param			beatId	path		int			true	"Beat ID"
//	@Param			beat	body		model.Beat	true	"Beat Data"
//	@Success		200		{string}	id
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/beatsheet/:id/beat/beatId  [put]
func PUTBeat(ctx *gin.Context) {
	var beat model.Beat
	id := ctx.Param("id")
	beatId := ctx.Param("beatId")
	mClient := data.NewMongoClient(ctx)
	ctx.ShouldBindJSON(&beat)
	res, err := mClient.UpdateBeat(id, beatId, beat)
	responseHandler(ctx, res, err)
}

// DELETEBeat godoc
//
//	@Summary		DELETE Beat
//	@Description	DELETE Beat by BeatId
//	@Tags			Beat
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int	true	"BeatSheet ID"
//	@Param			beatId	path		int	true	"Beat ID"
//	@Success		200		{string}	beatId
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/beatsheet/:id/beat/beatId  [delete]
func DELETEBeat(ctx *gin.Context) {
	id := ctx.Param("id")
	beatId := ctx.Param("beatId")
	mClient := data.NewMongoClient(ctx)
	err := mClient.DeleteBeat(id, beatId)
	responseHandler(ctx, nil, err)
}
