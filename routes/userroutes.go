package routes

import (
	"Chat_Api/database"
	"Chat_Api/models"

	"github.com/gofiber/fiber/v2"
)

// function to get a slice of all users
func GetUsers(c *fiber.Ctx) error {
	db := database.DBConn

	var users []models.User
	db.Find(&users)
	return c.JSON(users)
}

// function to get a specific user based on URL Params
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var user models.User
	db.Find(&user, id)
	return c.JSON(user)

}

// function to add a new user to the database (takes in "username" and "password" in JSON)
func NewUser(c *fiber.Ctx) error {
	db := database.DBConn
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&user)
	return c.JSON(user)
}

// function to delete a user with a specific id
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var user models.User
	db.First(&user, id)
	if user.Username == "" && user.Password == "" {
		return c.Status(500).SendString("No User Found with ID")
	}
	db.Delete(&user)
	return c.SendString("User Sucessfully Deleted")
}

// function to update a users password (takes "username" and "password" in JSON)
func UpdateUserPassword(c *fiber.Ctx) error {
	db := database.DBConn

	userInfo := new(models.UpdateUser)
	var user models.User

	if err := c.BodyParser(userInfo); err != nil {
		c.Status(500).SendString(" could not parse body of incoming JSON")
	}

	db.Raw("SELECT * FROM users WHERE username ='" + userInfo.Username + "'").Scan(&user)
	user.Password = userInfo.NewPassword
	db.Save(&user)

	return c.Status(204).SendString("The password for " + userInfo.Username + " has been successfully updated")

}
