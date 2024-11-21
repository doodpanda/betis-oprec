package main

import (
	"betis-oprec/database"
	"betis-oprec/router"

	"github.com/gofiber/fiber/v2"
)

// main is the entry point of the application. It initializes the Fiber app,
// connects to the database, performs database migrations, sets up routes,
// and starts the server on port 3000. If the server fails to start, it panics
// and logs the error.
func main() {
	app := fiber.New()
	database.ConnectDB()
	database.MigrateEnums(database.DB)
	database.MigrateDB()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("App is running properly! Served by Fiber")
	})

	router.SetupRoutes(app)
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
