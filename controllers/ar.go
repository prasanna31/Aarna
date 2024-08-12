package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prasanna31/Aarna/config"
	"github.com/prasanna31/Aarna/models"
)

func SendARMessage(c *gin.Context) {
	var arContent models.Message
	if err := c.ShouldBindJSON(&arContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	arContent.Type = "ar"
	if err := config.DB.Create(&arContent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "AR message sent"})
}

func GetARContent(c *gin.Context) {
	messageID := c.Param("messageId")
	var arContent models.Message
	if err := config.DB.Where("id = ? AND type = 'ar'", messageID).First(&arContent).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "AR content not found"})
		return
	}
	c.JSON(http.StatusOK, arContent)
}
