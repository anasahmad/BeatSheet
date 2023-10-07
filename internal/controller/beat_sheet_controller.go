package controller

import (
	"BeatSheet/internal/data"
	"BeatSheet/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type BeatSheetController struct {
	dataClient data.DataClient
}

func NewBeatSheetController(dataClient data.DataClient) BeatSheetController {
	return BeatSheetController{dataClient: dataClient}
}

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
func (c *BeatSheetController) POSTBeatSheet(ctx *gin.Context) {
	var beatSheet model.BeatSheet
	ctx.ShouldBindJSON(&beatSheet)
	id, err := c.dataClient.CreateBeatSheet(beatSheet)
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
func (c *BeatSheetController) GETBeatSheet(ctx *gin.Context) {
	id := ctx.Param("id")
	beatSheet, err := c.dataClient.RetrieveBeatSheet(id)
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
func (c *BeatSheetController) PUTBeatSheet(ctx *gin.Context) {
	var beatSheet model.BeatSheet
	id := ctx.Param("id")
	ctx.ShouldBindJSON(&beatSheet)
	res, err := c.dataClient.UpdateBeatSheet(id, &beatSheet)
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
func (c *BeatSheetController) DELETEBeatSheet(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.dataClient.DeleteBeatSheet(id)
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
func (c *BeatSheetController) GETBeatSheets(ctx *gin.Context) {
	baatSheets, err := c.dataClient.RetrieveBeatSheets()
	responseHandler(ctx, baatSheets, err)
}

func responseHandler(ctx *gin.Context, res interface{}, err interface{}) {
	if err != nil { // error within the app
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.(error).Error(),
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
