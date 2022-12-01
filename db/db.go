package db

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)
var DB *sqlx.DB

type Database struct {
	E *fiber.App
	Db *sqlx.DB
}
// var DSN = "user=vijay dbname=users sslmode=disable"
var DSN = "postgresql://vijay:12345@localhost:5432/users?sslmode=disable"
func ConnectDB() *Database {
	app := fiber.New()
	db, err := sqlx.Connect("postgres", DSN)
	if err != nil {
		fmt.Println(err)
		log.Fatalln("Connection to postgresql Failed")

	} 
	DB = db

	return &Database{
		Db: db,
		E: app,
	}
}
