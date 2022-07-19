package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marlonmp/wrixy/internal/apps/user/domain"
	"github.com/marlonmp/wrixy/internal/apps/user/port"
)

const (
	LocalsUserFilter = "userFilter"
	LocalsUserDomain = "userDomain"
)

type userServer struct {
	users port.UserService
}

func UserServer(s port.UserService) userServer {
	return userServer{s}
}

func (us userServer) List() fiber.Handler {

	return func(c *fiber.Ctx) error {

		var filter port.UserFilter

		if err := c.QueryParser(&filter); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		users, err := us.users.List(filter)

		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(users)
	}
}

func (us userServer) Post() fiber.Handler {

	return func(c *fiber.Ctx) error {

		user := c.Locals(LocalsUserDomain).(domain.User)

		createdUser, err := us.users.Post(user)

		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.Status(fiber.StatusCreated).JSON(createdUser)
	}
}

func (us userServer) Get() fiber.Handler {

	return func(c *fiber.Ctx) error {

		filter := c.Locals(LocalsUserFilter).(port.UserFilter)

		user, err := us.users.Get(filter)

		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(user)
	}
}

func (us userServer) Put() fiber.Handler {

	return func(c *fiber.Ctx) error {

		filter := c.Locals(LocalsUserFilter).(port.UserFilter)

		user := c.Locals(LocalsUserDomain).(domain.User)

		updatedUser, err := us.users.Update(filter.ID, user)

		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(updatedUser)
	}
}

// TODO: make patch route

func (us userServer) Delete() fiber.Handler {

	return func(c *fiber.Ctx) error {

		filter := c.Locals(LocalsUserFilter).(port.UserFilter)

		user, err := us.users.Delete(filter.ID)

		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(user)
	}
}
