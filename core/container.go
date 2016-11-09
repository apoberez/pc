package core

import (
	"database/sql"
	"fmt"
)

type Container struct {
	DB *sql.DB
}

func NewContainer() *Container {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", "oroinc", "oroinc", "pc")
	db, err := sql.Open("postgres", dbinfo)
	if nil != err {
		panic("Database connection")
	}
	return &Container{
		DB: db,
	}
}
