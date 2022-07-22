package fiber

import "github.com/gofiber/fiber/v2"

func (us userServer) Group(api fiber.Group) {

	users := api.Group("/users")

	users.Get("", UserQueryParser, us.list)

	users.Post("", us.post)

	usersByID := users.Group("/:uuid", UserParamsParser)

	usersByID.Get("", us.Get)

	usersByID.Put("", us.put)

	usersByID.Delete("", us.delete)
}
