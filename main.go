package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "118.98.64.74"
	port     = 5432
	user     = "postgres"
	password = "S@y@Br0nt0"
	dbname   = "belajar_golang"
)

type User struct {
	ID       string `json:"id" db="id"`
	Nama     string `json:"nama" db="username"`
	Password string `json:"password" db="password"`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	users := []User{}
	err = db.Select(&users, "SELECT * FROM user_eko")
	if err != nil {
		fmt.Println("dvdv")
	}
	fmt.Println("Successfully connected!")
}
