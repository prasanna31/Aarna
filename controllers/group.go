package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prasanna31/Aarna/config"
	"github.com/prasanna31/Aarna/models"
)

func CreateGroup(c *gin.Context) {
	var group models.Chat
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Group created"})
}

func AddMembersToGroup(c *gin.Context) {
	groupID := c.Param("groupId")
	var userIDs []uint
	if err := c.ShouldBindJSON(&userIDs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var group models.Chat
	if err := config.DB.Where("id = ?", groupID).First(&group).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
		return
	}
	for _, userID := range userIDs {
		var user models.User
		if err := config.DB.Where("id = ?", userID).First(&user).Error; err == nil {
			config.DB.Model(&group).Association("Members").Append(user)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Members added to group"})
}

func GetGroupDetails(c *gin.Context) {
	groupID := c.Param("groupId")
	var group models.Chat
	if err := config.DB.Preload("Members").Where("id = ?", groupID).First(&group).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
		return
	}
	c.JSON(http.StatusOK, group)
}

func RemoveMemberFromGroup(c *gin.Context) {
	groupID := c.Param("groupId")
	var userID uint
	if err := c.ShouldBindJSON(&userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var group models.Chat
	if err := config.DB.Where("id = ?", groupID).First(&group).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
		return
	}
	var user models.User
	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	config.DB.Model(&group).Association("Members").Delete(user)
	c.JSON(http.StatusOK, gin.H{"message": "Member removed from group"})
}
