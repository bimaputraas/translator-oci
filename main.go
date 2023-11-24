package main

import (
	"log"
	"translate/controller"
	"translate/logic"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := gin.Default()
	controllerBasic := controller.InitControllerBasic(logic.InitLogicBasic())
	controllerAdvance := controller.InitControllerAdvance(logic.InitLogicAdvance())
	app.GET("/basic", controllerBasic.Translate)
	app.GET("/advance", controllerAdvance.Translate)

	if err := app.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
