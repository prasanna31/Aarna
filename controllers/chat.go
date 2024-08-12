package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prasanna31/Aarna/config"
	"github.com/prasanna31/Aarna/models"
)

func SendMessage(c *gin.Context) {
	var msg models.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&msg).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Message sent"})
}

func GetMessages(c *gin.Context) {
	chatID := c.Param("chatId")
	var messages []models.Message
	if err := config.DB.Where("chat_id = ?", chatID).Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, messages)
}

func MarkMessagesRead(c *gin.Context) {
	var msgIDs []uint
	if err := c.ShouldBindJSON(&msgIDs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&models.Message{}).Where("id IN ?", msgIDs).Update("read", true).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Messages marked as read"})
}

func AcknowledgeMessage(c *gin.Context) {
	var msgID uint
	if err := c.ShouldBindJSON(&msgID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&models.Message{}).Where("id = ?", msgID).Update("read", true).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Message acknowledged"})
}
