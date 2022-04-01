package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kameshsampath/blogapp/graph"
	"github.com/kameshsampath/blogapp/graph/db"
	"github.com/kameshsampath/blogapp/graph/generated"
)

const defaultPort = "8080"

func main() {
	fmt.Println("Jai Guru")
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//Connect to DB
	d, err := db.ConnectToDB()
	defer d.Close()

	if err != nil {
		log.Fatalf("Error connecting to database %s, shutting down server.", err)
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: d}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
