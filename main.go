package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/eu-ovictor/gnlg/person"
	person_repository "github.com/eu-ovictor/gnlg/person/repository"
	person_usecase "github.com/eu-ovictor/gnlg/person/usecase"

	"github.com/eu-ovictor/gnlg/relationship"
	relationship_repository "github.com/eu-ovictor/gnlg/relationship/repository"
	relationship_usecase "github.com/eu-ovictor/gnlg/relationship/usecase"

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

	relationshipQuery := `
        CREATE TABLE IF NOT EXISTS relationship (
            ancestor INTEGER,
            descendant INTEGER,
            UNIQUE (ancestor, descendant)
        )
    `

	if _, err := db.ExecContext(ctx, relationshipQuery); err != nil {
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

	personRepository := person_repository.NewSQLitePersonRepository(db)
	personUsecase := person_usecase.NewPersonUsecase(personRepository)
	person.AddRoutes(router, personUsecase)

	relationshipRepository := relationship_repository.NewSQLiteRelationshipRepository(db)
	relationshipUsecase := relationship_usecase.NewRelationshipUsecase(relationshipRepository)
	relationship.AddRoutes(router, relationshipUsecase)

	if err := fasthttp.ListenAndServe(":42069", router.Handler); err != nil {
		panic(err)
	}
}
