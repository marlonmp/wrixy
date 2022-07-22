package fiber

import "github.com/gofiber/fiber/v2"

type Grouper interface {
	Group(fiber.Group)
}
