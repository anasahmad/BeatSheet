package controller

import (
	"BeatSheet/internal/data"
	"BeatSheet/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

// POSTBeatSheet godoc
//
//	@Summary		Post BeatSheet
//	@Description	Post BeatSheet from the object in the body
//	@Tags			BeatSheet
//	@Accept			json
//	@Produce		json
//	@Param			beatSheet	body		model.BeatSheet	true	"BeatSheet Data"
//	@Success		200			{string}	id
//	@Failure		400			{object}	error
//	@Failure		404			{object}	error
//	@Failure		500			{object}	error
//	@Router			/beatsheet  [post]
func POSTBeatSheet(ctx *gin.Context) {
	var beatSheet model.BeatSheet
	mClient := data.NewMongoClient(ctx)
	ctx.ShouldBindJSON(&beatSheet)
	id, err := mClient.CreateBeatSheet(beatSheet)
	responseHandler(ctx, id, err)
}

// GETBeatSheet godoc
//
//	@Summary		GET BeatSheet
//	@Description	GET BeatSheet by id
//	@Tags			BeatSheet
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"BeatSheet ID"
//	@Success		200	{string}	model.BeatSheet
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/beatsheet/:id  [get]
func GETBeatSheet(ctx *gin.Context) {
	id := ctx.Param("id")
	mClient := data.NewMongoClient(ctx)
	beatSheet, err := mClient.RetrieveBeatSheet(id)
	responseHandler(ctx, beatSheet, err)
}

// PUTBeatSheet godoc
//
//	@Summary		PUT BeatSheet
//	@Description	Update BeatSheet with the object in the body
//	@Tags			BeatSheet
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"BeatSheet ID"
//	@Param			BeatSheet	body		model.BeatSheet	true	"BeatSheet Data"
//	@Success		200			{string}	model.BeatSheet
//	@Failure		400			{object}	error
//	@Failure		404			{object}	error
//	@Failure		500			{object}	error
//	@Router			/beatsheet/:id  [put]
func PUTBeatSheet(ctx *gin.Context) {
	var beatSheet model.BeatSheet
	id := ctx.Param("id")
	mClient := data.NewMongoClient(ctx)
	ctx.ShouldBindJSON(&beatSheet)
	res, err := mClient.UpdateBeatSheet(id, &beatSheet)
	responseHandler(ctx, res, err)
}

// DELETEBeatSheet godoc
//
//	@Summary		DELETE BeatSheet
//	@Description	DELETE BeatSheet by id
//	@Tags			BeatSheet
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"BeatSheet ID"
//	@Success		200	{string}	id
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/beatsheet/:id  [delete]
func DELETEBeatSheet(ctx *gin.Context) {
	id := ctx.Param("id")
	mClient := data.NewMongoClient(ctx)
	err := mClient.DeleteBeatSheet(id)
	responseHandler(ctx, nil, err)
}

// GETBeatSheets godoc
//
//	@Summary		GET BeatSheets
//	@Description	GET All BeatSheets
//	@Tags			BeatSheet
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	[]model.BeatSheet
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/beatsheet  [get]
func GETBeatSheets(ctx *gin.Context) {
	mClient := data.NewMongoClient(ctx)
	baatSheets, err := mClient.RetrieveBeatSheets()
	responseHandler(ctx, baatSheets, err)
}

func responseHandler(ctx *gin.Context, res interface{}, err interface{}) {
	if err != nil { // error within the app
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	} else if res == nil { // delete
		ctx.AbortWithStatus(204)
	} else if reflect.TypeOf(res).String() == "string" { //added / created
		ctx.AbortWithStatusJSON(http.StatusCreated, gin.H{
			"id": res,
		})
	} else { // updated / retrieved
		ctx.JSON(http.StatusOK, res)
	}
}
