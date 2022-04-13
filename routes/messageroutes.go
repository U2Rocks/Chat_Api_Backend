package routes

import (
	"Chat_Api/database"
	"Chat_Api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// TODO: add many routes...

// function that grabs a message by the text contained within it (takes in "text" in JSON)
func GetMessageByText(c *fiber.Ctx) error {
	db := database.DBConn
	inputText := new(models.MessageSearch)
	if err := c.BodyParser(inputText); err != nil {
		c.Status(503).SendString(err.Error())
	}
	var messages []models.Message
	db.Raw("SELECT * FROM messages WHERE content LIKE '%" + inputText.Text + "%'").Scan(&messages)
	if len(messages) == 0 {
		c.Status(500).SendString("No messages found with the specified text")
	}
	return c.JSON(messages)
}

// function that gets a slice of all messages
func GetMessages(c *fiber.Ctx) error {
	db := database.DBConn

	var messages []models.Message
	db.Find(&messages)
	return c.JSON(messages)
}

// function that gets a single message via its ID
func GetMessage(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var message models.Message
	db.Find(&message, id)
	return c.JSON(message)
}

// function that adds a new message (takes in "content", "userid", "channelid" in JSON)
func NewMessage(c *fiber.Ctx) error {
	db := database.DBConn
	message := new(models.Message)

	if err := c.BodyParser(message); err != nil {
		c.Status(503).SendString(err.Error())
	}
	db.Create(&message)
	return c.JSON(message)
}

// function that deletes a message via ID
func DeleteMessage(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var message models.Message
	db.First(&message, id)
	if message.UserID == 0 {
		return c.Status(500).SendString("No message found with ID")
	}
	db.Delete(&message)
	return c.SendString("Message Sucessfully Deleted")
}

// function that gets all messages that are from a certain table (takes in "channelid" in JSON)
func GetMessagesByChatroomName(c *fiber.Ctx) error {
	db := database.DBConn

	// get chatroom id for chatroom that matches name input
	chatroomName := new(models.ChatroomSearch)
	var chatroom models.Chatroom

	if err := c.BodyParser(chatroomName); err != nil {
		c.Status(500).SendString("Json body could not be parsed: " + err.Error())
	}

	db.Raw("SELECT * FROM chatrooms WHERE title ='" + chatroomName.Chat_Name + "'").Scan(&chatroom)

	corrected_id := strconv.FormatUint(uint64(chatroom.ID), 10)

	// query messages table for all messages that share the id returned
	var messages []models.Message
	db.Raw("SELECT * FROM messages WHERE channel_id =" + corrected_id).Scan(&messages)

	return c.JSON(messages)
}
