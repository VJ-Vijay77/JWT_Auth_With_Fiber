package controllers

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error{
	return c.JSON("Welcome to Home")
}

func Register(c *fiber.Ctx)error{
return nil
}