package main

import (
	"log"
	database "track_doctor/db"

	"github.com/gofiber/fiber/v2"
	//"github.com/joho/godotenv"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
)

func printHello(c *fiber.Ctx) error{
	return c.SendString("Hello Baby")
}


func main() {
	database.ConnectDb()
	app := fiber.New()
	app.Get("/api", printHello)
	log.Fatal(app.Listen(":3000"))
}
