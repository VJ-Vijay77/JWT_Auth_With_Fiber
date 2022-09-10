package main

import (
	"github.com/VJ-Vijay77/JWT_Auth_With_Fiber/api"
	"github.com/VJ-Vijay77/JWT_Auth_With_Fiber/db"
	"github.com/VJ-Vijay77/JWT_Auth_With_Fiber/schema"
)



func main() {
	e := db.ConnectDB()
	e.Db.Exec(schema.User)	
	api.API(e.E)
	e.E.Listen(":8080")
}
