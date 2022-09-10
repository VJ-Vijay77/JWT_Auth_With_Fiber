package controllers

import (
	"github.com/VJ-Vijay77/JWT_Auth_With_Fiber/db"
	"github.com/VJ-Vijay77/JWT_Auth_With_Fiber/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Home(c *fiber.Ctx) error {
	return c.JSON("Welcome to Home")
}

func Register(c *fiber.Ctx) error {
	var user models.User
	c.BodyParser(&user)
	_, err := db.DB.Exec("INSERT INTO users(name,email,password)VALUES($1,$2,$3)", user.Name,user.Email,user.Password)
	if err != nil {
		return c.JSON("Failed")
	}
	return c.JSON(map[string]interface{}{
		"user details": user,
		"status":       "success",
	})
}


func Login(c *fiber.Ctx) error {
	var data models.User

	if err := c.BodyParser(&data); err != nil{
		return err
	}
	

	var user models.User
	 err := db.DB.Get(&user,"SELECT * FROM users WHERE email=$1",data.Email)
	if err != nil{
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message":"user not found",
		})
	}
	

	if data.Password != user.Password {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message":"wrong password",
		})
	}


	claims := jwt.NewWithClaims()

	return c.JSON(fiber.Map{
		"status":"Login Success",
		"user":user,
	})
}