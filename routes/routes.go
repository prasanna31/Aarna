package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prasanna31/Aarna/controllers"
)

func SetupRoutes(router *gin.Engine) {
	// Authentication routes
	router.POST("/api/register", controllers.Register)
	router.POST("/api/login", controllers.Login)
	router.POST("/api/logout", controllers.Logout)

	// User profile routes
	router.GET("/api/user/profile", controllers.GetProfile)
	router.PUT("/api/user/profile", controllers.UpdateProfile)

	// Chat routes
	router.POST("/api/messages/send", controllers.SendMessage)
	router.GET("/api/messages/:chatId", controllers.GetMessages)
	router.POST("/api/messages/read", controllers.MarkMessagesRead)
	router.POST("/api/messages/ack", controllers.AcknowledgeMessage)

	// AR messaging routes
	router.POST("/api/ar/send", controllers.SendARMessage)
	router.GET("/api/ar/:messageId", controllers.GetARContent)

	// Group chat routes
	router.POST("/api/groups/create", controllers.CreateGroup)
	router.PUT("/api/groups/:groupId/add", controllers.AddMembersToGroup)
	router.GET("/api/groups/:groupId", controllers.GetGroupDetails)
	router.DELETE("/api/groups/:groupId/remove", controllers.RemoveMemberFromGroup)

	// Media routes
	router.POST("/api/media/upload", controllers.UploadMedia)
	router.GET("/api/media/:mediaId", controllers.GetMedia)

	// Notification routes
	router.POST("/api/notifications/subscribe", controllers.SubscribeNotifications)
	router.POST("/api/notifications/unsubscribe", controllers.UnsubscribeNotifications)

	// Search routes
	router.GET("/api/search/users", controllers.SearchUsers)
	router.GET("/api/search/chats", controllers.SearchChats)

	// Settings routes
	router.GET("/api/settings", controllers.GetSettings)
	router.PUT("/api/settings", controllers.UpdateSettings)

	// Admin routes (if needed)
	router.GET("/api/admin/users", controllers.GetAllUsers)
	router.DELETE("/api/admin/user/:userId", controllers.DeleteUser)
}
