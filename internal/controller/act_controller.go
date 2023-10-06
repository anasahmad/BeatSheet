package controller

import (
	"BeatSheet/internal/ai"
	"BeatSheet/internal/data"
	"BeatSheet/internal/model"
	"github.com/gin-gonic/gin"
)

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
//	@Router			/beatsheet/:id/beat/:beatId/act  [post]
func POSTAct(ctx *gin.Context) {
	var act model.Act
	// Retrieving path params
	id := ctx.Param("id")
	beatId := ctx.Param("beatId")
	mClient := data.NewMongoClient(ctx)
	// Binding body to the struct
	ctx.ShouldBindJSON(&act)
	ai.NewAIModel().AddAct(act)
	actId, err := mClient.AddAct(id, beatId, act)
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
//	@Router			/beatsheet/:id/beat/:beatId/act/:actId  [put]
func PUTAct(ctx *gin.Context) {
	var act model.Act
	id := ctx.Param("id")
	beatId := ctx.Param("beatId")
	ActId := ctx.Param("actId")
	mClient := data.NewMongoClient(ctx)
	ctx.ShouldBindJSON(&act)
	res, err := mClient.UpdateAct(id, beatId, ActId, act)
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
//	@Router			/beatsheet/:id/beat/:beatId/act/:actId  [delete]
func DELETEAct(ctx *gin.Context) {
	id := ctx.Param("id")
	beatId := ctx.Param("beatId")
	ActId := ctx.Param("actId")
	mClient := data.NewMongoClient(ctx)
	err := mClient.DeleteAct(id, beatId, ActId)
	responseHandler(ctx, nil, err)
}
