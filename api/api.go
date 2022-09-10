package api

import (
	"github.com/VJ-Vijay77/JWT_Auth_With_Fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func API(e *fiber.App) {
	e.Get("/",controllers.Home)
	e.Post("/",controllers.Register)
}