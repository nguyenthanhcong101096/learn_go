package main

import (
	"fmt"
	"go_postgres/driver"
	repo "go_postgres/repository/repoiml"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "congttl"
	password = "password"
	dbname   = "wakuwaku_development"
)

func main() {
	db := driver.Connect(host, port, user, password, dbname)

	err := db.SQL.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("connect ok")

	repo := repo.NewUserRepo(db.SQL)
	arr, err := repo.Select()

	fmt.Println(arr)
}
