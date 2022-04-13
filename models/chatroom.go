package models

import (
	"gorm.io/gorm"
)

type Chatroom struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

type NewChatroomDescription struct {
	Title           string `json:"title"`
	New_Description string `json:"new_description"`
}

type NewChatroomTitle struct {
	Title     string `json:"title"`
	New_Title string `json:"new_title"`
}
