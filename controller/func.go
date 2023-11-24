package controller

import (
	"net/http"
	"translate/model"

	"github.com/gin-gonic/gin"
)

func (c *controllerBasic) Translate(ctx *gin.Context) {
	var req model.TranslateData
	req.SourceText = ctx.Query("text")
	req.TranslatedType = ctx.Query("target")

	data, raw, err := c.logic.Translate(req)
	if err != nil {
		ctx.JSON(500, model.Response{
			Message: err.Error(),
			Data:    raw,
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Message: "ok",
		Data:    data,
		Bix:     raw,
	})
}

func (c *controllerAdvance) Translate(ctx *gin.Context) {
	var req model.TranslateData
	req.SourceText = ctx.Query("text")
	req.TranslatedType = ctx.Query("target")
	// r.SourceType := c.Query("source")
	data, raw, err := c.logic.Translate(req)
	if err != nil {
		ctx.JSON(500, model.Response{
			Message: err.Error(),
			Bix:     raw,
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Message: "ok",
		Data:    data,
		Bix:     raw,
	})
}
