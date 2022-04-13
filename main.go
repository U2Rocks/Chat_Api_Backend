package main

// connect to database in new way on gorm.io/docs
// current database connection cannot read driver

import (
	"fmt"
	"log"

	"Chat_Api/database"
	"Chat_Api/models"
	"Chat_Api/routes"

	// "github.com/gofiber/fiber/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initializeRoutes(app *fiber.App) {
	app.Static("/", "./public/index.html") // make this route server a static html page...

	// routes that interact with the user table
	app.Get("/users", routes.GetUsers)
	app.Get("/users/:id", routes.GetUser)
	app.Post("/users", routes.NewUser)
	app.Delete("/users/:id", routes.DeleteUser)
	app.Post("/users/update", routes.UpdateUserPassword)

	// routes that interact with the chatroom table
	app.Get("/chatrooms", routes.GetChatrooms)
	app.Get("/chatrooms/:id", routes.GetChatroom)
	app.Post("/chatrooms", routes.NewChatroom)
	app.Delete("/chatrooms/:id", routes.DeleteChatroom)
	app.Post("/chatrooms/newtitle", routes.UpdateChatroomTitle)
	app.Post("/chatrooms/newdesc", routes.UpdateChatroomDescription)

	// routes that interact with messages
	app.Get("/messages", routes.GetMessages)
	app.Get("/messages/:id", routes.GetMessage)
	app.Post("/messages", routes.NewMessage)
	app.Post("/messages/textsearch", routes.GetMessageByText)
	app.Delete("/messages/:id", routes.DeleteMessage)
	app.Post("/messages/chatroom", routes.GetMessagesByChatroomName)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("chats.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	fmt.Println("Connection to Database now Open")
	database.DBConn.AutoMigrate(&models.User{}, &models.Chatroom{}, &models.Message{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New() // create a new fiber app instance

	// app.Use(middleware.Logger()) // add a logger to see incoming requests

	// initialize database and defer the connection closing
	initDatabase()
	sqlDB, err := database.DBConn.DB()
	if err != nil {
		panic("failed to create generic db object: " + err.Error())
	}
	defer sqlDB.Close()

	initializeRoutes(app) // initialize routes in separate function for readability

	log.Fatal(app.Listen(":3000"))

}
