package main

import (
	"context"
	"database/sql"
	"log"

	"sqlc-tutorial/db/sqlc"

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

	accountParams := sqlc.CreateAccountParams{
		Owner:    "Jorge Numa",
		Balance:  1000,
		Currency: "COP",
	}

	// create an account
	insertedAccount, err := queries.CreateAccount(ctx, accountParams)

	if err != nil {
		return err
	}
	log.Println("Inserted", insertedAccount)

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
