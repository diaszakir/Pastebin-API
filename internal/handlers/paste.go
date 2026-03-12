package handlers

import (
	"net/http"
	"time"

	"github.com/diaszakir/base62-go"
	"github.com/diaszakir/gopastebin/internal/models"
	"github.com/gin-gonic/gin"
)

const URL = "http://localhost:8080"

var nextID int64 = 0

var pasteInfo = make(map[string]models.Paste)

func CreatePaste(c *gin.Context) {
	var paste models.PasteRequest
	err := c.ShouldBindJSON(&paste)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	nextID++
	code, err := base62.Encode(nextID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	url := models.PasteResponse{
		URL: URL + "/" + code,
	}

	pasteInfo[code] = models.Paste{
		ID:        nextID,
		Code:      code,
		Content:   paste.Content,
		CreatedAt: time.Now(),
	}

	c.JSON(http.StatusCreated, gin.H{
		"url":  url.URL,
		"code": code,
	})
}

func GetPaste(c *gin.Context) {
	code := c.Param("code")

	paste, ok := pasteInfo[code]

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "paste not found"})
		return
	}

	c.JSON(http.StatusOK, paste)
}

func DeletePaste(c *gin.Context) {
	code := c.Param("code")

	_, ok := pasteInfo[code]

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "paste not found"})
		return
	}

	delete(pasteInfo, code)

	c.Status(http.StatusNoContent)
}

func RawPaste(c *gin.Context) {
	code := c.Param("code")

	paste, ok := pasteInfo[code]

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "paste not found"})
		return
	}

	c.String(http.StatusOK, paste.Content)
}
