// Copyright 2019-20 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package db_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sqlc-tutorial/db/sqlc"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	ctx = context.Background()
)

// Add DB credentials here
var (
	user   = "jnuma"
	host   = "localhost"
	port   = "5432"
	dbname = "bank"
)

func init() {
	paramsDB := fmt.Sprintf("user=%s dbname=%s sslmode=disable", user, dbname)
	db, err := sql.Open("postgres", paramsDB)
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Paila con ese PING")
		panic(err)
	}
}

type accountModel struct {
	ID       int    `gorm:"column:id" db:"id"`
	Owner    string `gorm:"column:owner" db:"owner"`
	Balance  int    `gorm:"column:balance" db:"balance"`
	Currency string `gorm:"column:currency" db:"currency"`
}

// Required by gorm
func (accountModel) TableName() string {
	return "account"
}

func Benchmark(b *testing.B) {
	setup()
	defer cleanup()

	limits := []int{
		5,
		50,
		500,
		10000,
	}

	paramsDB := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable", host, user, dbname, port)

	for _, lim := range limits {

		// Benchmark sqlc
		db, _ := sql.Open("postgres", paramsDB)
		queries := sqlc.New(db)
		b.Run(fmt.Sprintf("sqlc limit:%d", lim), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				accounts, err := queries.ListAccountLimit(ctx, int32(lim))
				if err != nil {
					b.Fatal(err)
				}

				if len(accounts) != lim {
					panic("something is wrong")
				}
			}
		})

		// Benchmark gorm
		gormDB, err := gorm.Open("postgres", paramsDB)
		if err != nil {
			panic(err)
		}

		b.Run(fmt.Sprintf("gorm limit:%d", lim), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var res = []accountModel{}

				err := gormDB.Order("owner").Limit(lim).Find(&res).Error
				if err != nil {
					panic(err)
				}
				if len(res) != lim {
					panic("something is wrong")
				}
			}
		})

		fmt.Println("========================================================================")
	}

}

func setup() {
	ctx := context.Background()
	db, err := sql.Open("postgres", "user=jnuma dbname=bank sslmode=disable")
	if err != nil {
		panic(err)
	}

	queries := sqlc.New(db)

	// Add 10,000 fake account entries
	for i := 0; i < 10000; i++ {
		accountParams := sqlc.CreateAccountParams{
			Owner:    gofakeit.Name(),
			Balance:  gofakeit.Int64(),
			Currency: "COP",
		}
		_, err = queries.CreateAccount(ctx, accountParams)
		if err != nil {
			log.Fatalln("PAILA ", err.Error())
		}
	}
}

func cleanup() {
	// run here migration down
}
