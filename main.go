package main

import (
	"github.com/VJ-Vijay77/JWT_Auth_With_Fiber/api"
	"github.com/VJ-Vijay77/JWT_Auth_With_Fiber/db"
	"github.com/VJ-Vijay77/JWT_Auth_With_Fiber/schema"
	"github.com/gofiber/fiber/v2/middleware/cors"
)



func main() {
	e := db.ConnectDB()
	e.Db.Exec(schema.User)	
	e.E.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	api.API(e.E)
	e.E.Listen(":8080")
}
