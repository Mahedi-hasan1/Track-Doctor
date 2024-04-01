package main

import (
	"log"
	database "track_doctor/db"
	"track_doctor/routes"

	"github.com/gofiber/fiber/v2"
)



func main() {
	database.ConnectDb()
	//database.AutoMigrate(database.DB.Db)
	app := fiber.New()
	routes.Setup(app)
	log.Fatal(app.Listen(":3000"))
}
