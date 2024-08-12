package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ChatID    uint
	SenderID  uint
	Content   string
	Type      string // "text", "image", "ar", "video"
	Read      bool
	Timestamp int64
}
