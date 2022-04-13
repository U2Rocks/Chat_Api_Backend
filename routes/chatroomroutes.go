package routes

import (
	"Chat_Api/database"
	"Chat_Api/models"

	"github.com/gofiber/fiber/v2"
)

// TODO: add functions that add a new chatroom and list all chatrooms

// function to list all chatrooms
func GetChatrooms(c *fiber.Ctx) error {
	db := database.DBConn
	var chatrooms []models.Chatroom

	db.Find(&chatrooms)
	return c.JSON(chatrooms)
}

// function to get a specific chatrooms info
func GetChatroom(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var chatroom models.Chatroom
	db.Find(&chatroom, id)
	return c.JSON(chatroom)
}

// function to add a new chatroom to the database (takes in "title" and "description" in JSON)
func NewChatroom(c *fiber.Ctx) error {
	db := database.DBConn
	chatroom := new(models.Chatroom)

	if err := c.BodyParser(chatroom); err != nil {
		c.Status(503).SendString(err.Error())
	}
	db.Create(&chatroom)
	return c.JSON(chatroom)
}

// function to delete a chatrooms with a specific id
func DeleteChatroom(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var chatroom models.Chatroom
	db.First(&chatroom, id)
	if chatroom.Title == "" && chatroom.Description == "" {
		return c.Status(500).SendString("No User Found with ID")
	}
	db.Delete(&chatroom)
	return c.SendString("Chatroom Sucessfully Deleted")
}

// function to update a chatrooms title (takes "title" and "new_title" in JSON)
func UpdateChatroomTitle(c *fiber.Ctx) error {
	db := database.DBConn

	chatroomInfo := new(models.NewChatroomTitle)
	var chatroom models.Chatroom

	if err := c.BodyParser(chatroomInfo); err != nil {
		c.Status(500).SendString("Could not parse JSON of request")
	}

	db.Raw("SELECT * FROM chatrooms WHERE title = '" + chatroomInfo.Title + "'").Scan(&chatroom)
	chatroom.Title = chatroomInfo.New_Title
	db.Save(&chatroom)

	return c.Status(204).SendString("entry " + chatroomInfo.Title + " successfully updated to: " + chatroomInfo.New_Title)
}

// function to update a chatrooms description (takes "title" and "new_description" in JSON)
func UpdateChatroomDescription(c *fiber.Ctx) error {
	db := database.DBConn

	chatroomInfo := new(models.NewChatroomDescription)
	var chatroom models.Chatroom

	if err := c.BodyParser(chatroomInfo); err != nil {
		c.Status(500).SendString("Could not parse JSON of request")
	}

	db.Raw("SELECT * FROM chatrooms WHERE title ='" + chatroomInfo.Title + "'").Scan(&chatroom)
	chatroom.Description = chatroomInfo.New_Description
	db.Save(&chatroom)

	return c.Status(204).SendString("entry " + chatroom.Title + " successfully updated description to: " + chatroomInfo.New_Description)
}
