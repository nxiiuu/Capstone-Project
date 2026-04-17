package entity

import (
	"gorm.io/gorm"
)

type Chatrooms struct {
	gorm.Model
	ChatRoomID string `gorm:"primaryKey" json:"chat_room_id"`
	GroupName  string `json:"group_name"`
	RoomType   string `json:"room_type"`
	SGroupID   *uint  `json:"s_group_id"`

	Messages []Messages `gorm:"foreignKey:ChatRoomID" json:"messages"`
}
