package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadMedia(c *gin.Context) {
	file, _ := c.FormFile("media")
	filePath := "./uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload media"})
		return
	}
	// Save media details to database
	c.JSON(http.StatusOK, gin.H{"message": "Media uploaded", "path": filePath})
}

func GetMedia(c *gin.Context) {
	mediaID := c.Param("mediaId")
	// Retrieve media path from database
	mediaPath := "./uploads/" + mediaID // Example path
	c.File(mediaPath)
}
