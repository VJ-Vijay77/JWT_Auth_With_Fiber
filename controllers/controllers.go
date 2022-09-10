package controllers

import (
	"github.com/VJ-Vijay77/JWT_Auth_With_Fiber/db"
	"github.com/VJ-Vijay77/JWT_Auth_With_Fiber/models"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.JSON("Welcome to Home")
}

func Register(c *fiber.Ctx) error {
	var user models.User
	c.BodyParser(&user)
	_, err := db.DB.Exec("INSERT INTO users(name,email,password)VALUES($1,$2,$3)", "VIJAY", "vijay@gmai.com", "1")
	if err != nil {
		return c.JSON("Failed")
	}
	return c.JSON(map[string]interface{}{
		"user details": user,
		"status":       "success",
	})
}


func Login(c *fiber.Ctx) error {
	return nil
}