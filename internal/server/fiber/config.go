package fiber

import (
	"github.com/gofiber/fiber/v2"
	json "github.com/json-iterator/go"
)

var config fiber.Config

func init() {
	config = fiber.Config{
		StrictRouting: true,
		CaseSensitive: true,
		AppName:       "wrixy",
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	}
}
