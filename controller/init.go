package controller

import (
	"translate/logic"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	Translate(ctx *gin.Context)
}

type controllerBasic struct {
	logic logic.Logic
}

type controllerAdvance struct {
	logic logic.Logic
}

func InitControllerBasic(l logic.Logic) Controller {
	return &controllerBasic{logic: l}
}

func InitControllerAdvance(l logic.Logic) Controller {
	return &controllerAdvance{logic: l}
}
