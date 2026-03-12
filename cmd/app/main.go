package main

import (
	"github.com/diaszakir/gopastebin/internal/routes"
	"github.com/gin-gonic/gin"

	_ "github.com/diaszakir/gopastebin/docs"
)

// @Title Pastebin API
// @Description Pastebin API for sending messages through link
// @Version 1.0
// @Host localhost:8080
// @BasePath /
func main() {
	app := gin.Default()

	routes.Routes(app)

	app.Run(":8080")
}
