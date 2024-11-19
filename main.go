package main

import (
	"betis-oprec/database"
	"betis-oprec/router"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.ConnectDB()
	database.MigrateEnums(database.DB)
	fmt.Println("ENUMS MIGRATED")
	database.MigrateDB()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("App is running properly! Served by Fiber")
	})

	router.SetupRoutes(app)
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
