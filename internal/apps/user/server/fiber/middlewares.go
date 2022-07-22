package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marlonmp/wrixy/internal/apps/user/domain"
	"github.com/marlonmp/wrixy/internal/apps/user/port"
)

func UserDomainParser(c *fiber.Ctx) error {

	var user domain.User

	if err := c.BodyParser(&user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	c.Locals(LocalsUserDomain, user)

	return c.Next()
}

func UserQueryParser(c *fiber.Ctx) error {

	var filter port.UserFilter

	if err := c.QueryParser(&filter); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	c.Locals(LocalsUserFilter, filter)

	return c.Next()
}

func UserParamsParser(c *fiber.Ctx) error {

	var filter port.UserFilter

	if err := c.ParamsParser(&filter); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	c.Locals(LocalsUserFilter, filter)

	return c.Next()
}
