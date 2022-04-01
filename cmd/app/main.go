package main

import (
	"fmt"
	"net/http"
	"os"
	"path"

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

	//Intialize the Database
	dbFile, ok := os.LookupEnv("TODOS_DB_FILE")
	if !ok {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error getting current working directory %s", err)
		}

		err = os.MkdirAll(path.Join(cwd, "work"), os.ModeDir)
		if err != nil {
			log.Fatalf("Error making db directory %s", err)
		}

		dbFile = fmt.Sprintf("%s/todo.db", path.Join(cwd, "work"))
	}

	d, err := db.InitDB(dbFile)
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
