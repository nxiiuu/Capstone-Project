package entity

import (
	"gorm.io/gorm"
)

type Messages struct {
	gorm.Model
	Body           string `json:"body"`
	MessagesTypeID *uint  `json:"messages_type_id"`
	UserID         *uint  `json:"user_id"`
	ChatRoomID     string `json:"chat_room_id"`

	Chatroom     *Chatrooms    `gorm:"foreignKey:ChatRoomID" json:"chatroom"`
	MessagesType *MessagesType `gorm:"foreignKey:MessagesTypeID" json:"messages_type"`
	User         *User         `gorm:"foreignKey:UserID" json:"user"`
}
