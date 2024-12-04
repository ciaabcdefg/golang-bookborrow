package main

import (
	"borrow/internal/db"
	"borrow/internal/env"
	"borrow/internal/prefixedrouter"
	"borrow/repo"
	"borrow/services/auth"
	"borrow/services/books"
	"borrow/services/borrow"
	"borrow/services/students"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Hello world"))
}

func main() {
	router := httprouter.New()
	subrouter := prefixedrouter.PrefixedRouter{
		Router: router,
		Prefix: "/api/v1",
	}

	subrouter.GET("/", Index)

	db, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	repo := repo.New(db)

	books.NewHandler(repo, db).RegisterRoutes(subrouter)
	students.NewHandler(repo, db).RegisterRoutes(subrouter)
	borrow.NewHandler(repo, db).RegisterRoutes(subrouter)
	auth.NewHandler(repo, db).RegisterRoutes(subrouter)

	log.Println("Connected to the database")
	log.Printf("Listening on %s:%s\n", env.Hostname, env.Port)

	corsHandler := cors.Default().Handler(router)
	http.ListenAndServe(env.Hostname+":"+env.Port, corsHandler)
}
