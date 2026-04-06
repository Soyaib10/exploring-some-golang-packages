package main

import (
	"context"
	"fmt"
	"log"

	db "github.com/Soyaib10/exploring-some-golang-packages/sqlc/db/generated"
	"github.com/jackc/pgx/v5"
)

func main() {
	connStr := "postgres://postgres:postgres@localhost/sqlc"

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	queries := db.New(conn)

	user, err := queries.GetUserByEmail(context.Background(), "test@gmail.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)

	err = queries.DeleteUser(context.Background(), 4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted!")

	users, err := queries.GetAllUsers(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range users {
		fmt.Println(u)
	}

	tx, err := conn.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback(context.Background())

	qtx := queries.WithTx(tx)

	newUser, err := qtx.CreateUser(context.Background(), db.CreateUserParams{
		Username: "salam",
		Email:    "salam@gmail.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newUser)

	err = tx.Commit(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("committed!")
}
