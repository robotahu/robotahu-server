package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"robotahu-server/gql"
	"robotahu-server/postgres"
	"robotahu-server/server"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
)

func main() {
	router := chi.NewRouter()

	db, err := postgres.New(
		postgres.ConnString(
			os.Getenv("POSTGRE_HOST"),
			os.Getenv("POSTGRE_USER"),
			os.Getenv("POSTGRE_PASSWORD"),
			os.Getenv("POSTGRE_DB_NAME"),
		),
	)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	rootQuery := gql.NewRoot(db)

	sc, err := graphql.NewSchema(graphql.SchemaConfig{Query: rootQuery.Query})
	if err != nil {
		fmt.Println("Error creating schema:", err)
	}

	s := server.Server{GqlSchema: &sc}

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.Compress(5),
		middleware.StripSlashes,
		middleware.Recoverer,
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIO    NS"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of maj    or browsers
		}),
	)

	router.Post("/graphql", s.GraphQL())

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
