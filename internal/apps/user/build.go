package user

import (
	psqlRepo "github.com/marlonmp/wrixy/internal/apps/user/repo/psql"
	"github.com/marlonmp/wrixy/internal/apps/user/server/fiber"
	"github.com/marlonmp/wrixy/internal/apps/user/service"
	"github.com/marlonmp/wrixy/internal/repo/psql"
	fiberServer "github.com/marlonmp/wrixy/internal/server/fiber"
)

func BuildServer() fiberServer.Grouper {

	client := psql.Client()

	repo := psqlRepo.UserRepo(client)

	service := service.UserService(repo)

	return fiber.UserServer(service)
}
