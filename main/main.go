package main

import (
	"github.com/go-martini/martini"
	"../handlers"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var dbError error


func main() {

	db, dbError = sql.Open("sqlite3", "./db.sqlite")
	if dbError != nil { panic(dbError) }
	if db == nil { panic("db nil") }

	m := martini.Classic()
	m.Map(db)
	// m.Use(func(res http.ResponseWriter) {
	// 	res.Header().Set("Content-Type", "application/json")
	// })

	m.Group(`/auth`, func(r martini.Router) {
		r.Get(`/profile`, handlers.HandleProfile)
		r.Get(`/login`, handlers.HandleLogin)
	})

	m.Get(`/user`, handlers.HandleWs)

	m.RunOnAddr(":3007")
}
