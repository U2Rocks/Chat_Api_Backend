package models

// TODO: add refer types to manage relationships with other tables

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Content   string `json:"content"`
	UserID    uint   `json:"userid"`
	ChannelID uint   `json:"channelid"`
}

type MessageSearch struct {
	Text string `json:"text"`
}

type ChatroomSearch struct {
	Chat_Name string `json:"chat_name"`
}
