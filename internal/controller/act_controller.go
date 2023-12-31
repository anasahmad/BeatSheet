package controller

import (
	"BeatSheet/internal/ai"
	"BeatSheet/internal/data"
	"BeatSheet/internal/model"
	"github.com/gin-gonic/gin"
)

type ActController struct {
	dataClient data.DataClient
	aiModel    *ai.AIModel
}

func NewActController(dataClient data.DataClient, aiModel *ai.AIModel) ActController {
	return ActController{dataClient: dataClient, aiModel: aiModel}
}

// POSTAct godoc
//
//	@Summary		Post Act
//	@Description	Post Act from the object in the body
//	@Tags			Act
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"BeatSheet ID"
//	@Param			beatId	path		int			true	"Beat ID"
//	@Param			act		body		model.Act	true	"Act Data"
//	@Success		200		{string}	id
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/:id/beat/:beatId/act  [post]
func (c *ActController) POSTAct(ctx *gin.Context) {
	var act model.Act
	// Retrieving path params
	id := ctx.Param("id")
	beatId := ctx.Param("beatId")
	// Binding body to the struct
	ctx.ShouldBindJSON(&act)
	c.aiModel.AddAct(act)
	actId, err := c.dataClient.AddAct(id, beatId, act)
	responseHandler(ctx, actId, err)
}

// PUTAct godoc
//
//	@Summary		PUT Act
//	@Description	Update Act from the object in the body
//	@Tags			Act
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"BeatSheet ID"
//	@Param			beatId	path		int			true	"Beat ID"
//	@Param			actId	path		int			true	"Act ID"
//	@Param			act		body		model.Act	true	"Act Data"
//	@Success		200		{string}	id
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/:id/beat/:beatId/act/:actId  [put]
func (c *ActController) PUTAct(ctx *gin.Context) {
	var act model.Act
	id := ctx.Param("id")
	beatId := ctx.Param("beatId")
	ActId := ctx.Param("actId")
	ctx.ShouldBindJSON(&act)
	res, err := c.dataClient.UpdateAct(id, beatId, ActId, act)
	responseHandler(ctx, res, err)
}

// DELETEAct godoc
//
//	@Summary		DELETE Act
//	@Description	Delete Act by actId
//	@Tags			Act
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int	true	"BeatSheet ID"
//	@Param			beatId	path		int	true	"Beat ID"
//	@Param			actId	path		int	true	"Act ID"
//	@Success		200		{string}	id
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/:id/beat/:beatId/act/:actId  [delete]
func (c *ActController) DELETEAct(ctx *gin.Context) {
	id := ctx.Param("id")
	beatId := ctx.Param("beatId")
	ActId := ctx.Param("actId")
	err := c.dataClient.DeleteAct(id, beatId, ActId)
	responseHandler(ctx, nil, err)
}
