package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/yurirennan/graphql-with-go/graph"
	"github.com/yurirennan/graphql-with-go/internal/database"
)

const defaultPort = "8080"

func main() {
	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	defer db.Close()

	categoryDB := database.NewCategory(db)
	courseDB := database.NewCourse(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CategoryDB: categoryDB,
		CourseDB:   courseDB,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
