package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/health", checkAPI)
}

func checkAPI(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API response"})
}
