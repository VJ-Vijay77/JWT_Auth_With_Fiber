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

	t := jwt.NewNumericDate(time.Now().Add(time.Minute *3))

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.RegisteredClaims{
		Issuer: strconv.Itoa(int(user.Id)),
		ExpiresAt: t,
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

func User (c *fiber.Ctx) error {
	cookie :=  c.Cookies("jwt")
	token,err := jwt.ParseWithClaims(cookie,&jwt.RegisteredClaims{},func(t *jwt.Token) (interface{}, error) {
		return []byte(JWTsecret),nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message":"Unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.RegisteredClaims)
	var user models.User
	db.DB.Get(&user,"SELECT * FROM users WHERE id=$1",claims.Issuer)

	return c.JSON(user)
}


func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message":"Logout Success",
	})
}