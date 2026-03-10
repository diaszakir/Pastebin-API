package routes

import (
	"net/http"

	"github.com/diaszakir/gopastebin/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/health", checkAPI)

	r.POST("/paste", handlers.CreatePaste)

	r.GET("/paste/:id")

	r.DELETE("/paste/:id")
}

func checkAPI(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API response"})
}
