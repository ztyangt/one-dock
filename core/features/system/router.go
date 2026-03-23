package system

import (
	"github.com/gofiber/fiber/v3"
)

func Setup(routerGroup fiber.Router) {

	systems := routerGroup.Group("/system")
	systems.Get("/", systemInfo)
}
