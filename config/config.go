package config

import "github.com/jmoiron/sqlx"

type Values struct {
	Db *sqlx.DB
}
