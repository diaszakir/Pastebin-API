package routes

import (
	"net/http"

	"github.com/diaszakir/gopastebin/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/health", checkAPI)

	r.POST("/paste", handlers.CreatePaste)

	r.GET("/paste/:code", handlers.GetPaste)

	r.DELETE("/paste/:code", handlers.DeletePaste)
}

func checkAPI(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API response"})
}
