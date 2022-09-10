package controllers

import (
	"strconv"
	"time"

	"github.com/VJ-Vijay77/JWT_Auth_With_Fiber/db"
	"github.com/VJ-Vijay77/JWT_Auth_With_Fiber/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const JWTsecret = "secret"

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


	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Minute *2).Unix(),
	})

	token,err := claims.SignedString([]byte(JWTsecret))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message":"could not login",
		})
	}

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Minute *2),
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"status":"Login Success",
		"token":token,
	})
}