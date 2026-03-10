package main

import (
	"github.com/diaszakir/gopastebin/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	routes.Routes(app)

	app.Run(":8080")
}
