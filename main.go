package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/eu-ovictor/gnlg/person"
	"github.com/eu-ovictor/gnlg/person/repository"
	"github.com/eu-ovictor/gnlg/person/usecase"
	"github.com/fasthttp/router"
	_ "github.com/mattn/go-sqlite3"
	"github.com/valyala/fasthttp"
)

func setup(ctx context.Context, db *sql.DB) error {
	personQuery := `
        CREATE TABLE IF NOT EXISTS person (
            rowid INTEGER PRIMARY KEY,
            name TEXT NOT NULL 
        )
    `

	if _, err := db.ExecContext(ctx, personQuery); err != nil {
		return fmt.Errorf("error creating person table: %w", err)
	}

	return nil
}

func main() {
	db, err := sql.Open("sqlite3", "gnlg.db")
	if err != nil {
		panic(err)
	}

	ctx, cancelFunc := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancelFunc()

	if err := setup(ctx, db); err != nil {
		panic(err)
	}

	router := router.New()

    router.GET("/", func(ctx *fasthttp.RequestCtx) {
        ctx.SetStatusCode(200)
        ctx.SetBodyString("I'm alive!")
    })

	personRepository := repository.NewSQLitePersonRepository(db)
	personUsecase := usecase.NewPersonUsecase(personRepository)
	person.AddRoutes(router, personUsecase)

	if err := fasthttp.ListenAndServe(":42069", router.Handler); err != nil {
		panic(err)
	}
}
