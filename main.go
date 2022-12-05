package main

import (
	"context"
	"database/sql"
	"log"

	"sqlc-tutorial/sqlc"

	_ "github.com/lib/pq"
)

func run() error {
	ctx := context.Background()

	db, err := sql.Open("postgres", "user=jnuma dbname=bank sslmode=disable")
	if err != nil {
		return err
	}

	queries := sqlc.New(db)

	// list all accounts
	accounts, err := queries.ListAccount(ctx)
	if err != nil {
		return err
	}
	log.Println(accounts)

	// create an account

}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
