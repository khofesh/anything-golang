package main

import "database/sql"

// Post ..
type Post struct {
	ID      int
	Content string
	Author  string
}

// Db ...
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}
