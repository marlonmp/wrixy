package fiber

import "github.com/gofiber/fiber/v2"

func routesUp(api fiber.Group) {

	users := api.Group("/users")

	users.Get("")

	users.Post("")

	usersByID := users.Group("/:uuid")

	usersByID.Get("")

	usersByID.Put("")

	usersByID.Patch("")

	usersByID.Delete("")
}
