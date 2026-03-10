package handlers

import (
	"net/http"
	"time"

	"github.com/diaszakir/base62-go"
	"github.com/diaszakir/gopastebin/internal/models"
	"github.com/gin-gonic/gin"
)

const URL = "http://localhost:8080"

var nextID = 0

var pasteInfo = make(map[int]models.Paste)

func CreatePaste(c *gin.Context) {
	var paste models.PasteRequest
	err := c.ShouldBindJSON(&paste)

	if err != nil {
		return
	}

	nextID++
	code, err := base62.Encode(int64(nextID))

	url := models.PasteResponse{
		URL: URL + "/" + code,
	}

	pasteInfo[nextID] = models.Paste{
		ID:        int64(nextID),
		Content:   paste.Content,
		CreatedAt: time.Now(),
	}

	c.JSON(http.StatusOK, url)
}
