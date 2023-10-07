package controller

import (
	"BeatSheet/internal/ai"
	"BeatSheet/internal/data"
	"BeatSheet/internal/model"
	"github.com/gin-gonic/gin"
)

type BeatController struct {
	dataClient data.DataClient
	aiModel    *ai.AIModel
}

func NewBeatController(dataClient data.DataClient, aiModel *ai.AIModel) BeatController {
	return BeatController{dataClient: dataClient, aiModel: aiModel}
}

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
//	@Router			/:id/beat  [post]
func (c *BeatController) POSTBeat(ctx *gin.Context) {
	var beat model.Beat
	id := ctx.Param("id")
	ctx.ShouldBindJSON(&beat)
	ai.NewAIModel().AddBeat(beat)
	beatId, err := c.dataClient.AddBeat(id, beat)
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
//	@Router			/:id/beat/beatId  [put]
func (c *BeatController) PUTBeat(ctx *gin.Context) {
	var beat model.Beat
	id := ctx.Param("id")
	beatId := ctx.Param("beatId")
	ctx.ShouldBindJSON(&beat)
	res, err := c.dataClient.UpdateBeat(id, beatId, beat)
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
//	@Router			/:id/beat/beatId  [delete]
func (c *BeatController) DELETEBeat(ctx *gin.Context) {
	id := ctx.Param("id")
	beatId := ctx.Param("beatId")
	err := c.dataClient.DeleteBeat(id, beatId)
	responseHandler(ctx, nil, err)
}
