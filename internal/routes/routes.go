package routes

import (
	"net/http"

	"github.com/diaszakir/gopastebin/internal/handlers"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes(r *gin.Engine) {
	r.GET("/health", checkAPI)
	r.GET("/paste/:code", handlers.GetPaste)
	r.GET("/paste/raw/:code", handlers.RawPaste)

	r.POST("/paste", handlers.CreatePaste)

	r.DELETE("/paste/:code", handlers.DeletePaste)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// CheckAPI godoc
// @Summary Check REST API
// @Description Checking API for working
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func checkAPI(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API response"})
}
