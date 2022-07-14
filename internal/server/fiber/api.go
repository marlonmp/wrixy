package fiber

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func init() {

	hostPort := os.Getenv("HOST_PORT")

	app := fiber.New(config)

	api := app.Group("/api")

	routesUp(api)

	app.Listen(hostPort)
}
